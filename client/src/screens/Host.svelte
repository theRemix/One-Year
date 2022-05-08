<script>
  const HostState = {
    CREATE_GAME: 0,
    PROMPT: 1,
    RESOLVE: 2,
  }

  let hostState = HostState.CREATE_GAME

  let code = localStorage.getItem('code')

  let errMessage = ""

  // ==================== State =====================

  const playerNamesSet = new Set()
  $: playerNames = []
  const addPlayerName = playerName => {
    playerNamesSet.add(playerName)
    playerNames = [...playerNamesSet]

    if (!playerScores.hasOwnProperty(playerName)) {
      playerScores = {
        ...playerScores,
        [playerName]: 0,
      }
    }
  }

  const playerAnswers = new Map()
  $: playersAnswered = []
  const setPlayerAnswer = ({playerName, answers}) => {
    // save answer
    playerAnswers.set(playerName, answers)

    // mark as answered
    playersAnswered = [...playerAnswers.keys()]
  }

  // ====================== SSE =====================
  
  const startSSE = () => {

    const events = new EventSource(`/events/host/${code}`)

    // sse needs plain text
    events.onmessage = function(event) {
      const payload = JSON.parse(event.data)
      const { op } = payload

      switch (op) {
        case proto.PlayerEvent.Op.READY:
          if (payload.name === name) {
            hostState = HostState.PROMPT
          } 
          addPlayerName(payload.playerName)
          break
        case proto.PlayerEvent.Op.ANSWERED:
          setPlayerAnswer(payload.answer)
          break
        default:
          console.warn(`SSE received unsupported Op: ${op}`)
      }
    }

    events.onerror = function(error) {
      console.error('onerror sse', error)
      errMessage = error.message || ''
      setTimeout(() => errMessage = '', 3000)
    }
  }

  // ============ Restore Ongoing Game =================

  if (code) {
    hostState = HostState.PROMPT
    startSSE()
  }

  // ====================== Create =====================

  const createGame = () => {
    const msg = new proto.Join(); // reuse

    msg.setCode(code);

    const body = msg.serializeBinary()

    fetch('/api/host/create', {
      method: 'POST',
      'content-type': 'application/x-protobuf',
      body,
    })
    .then(response => response.arrayBuffer())
    .then(data => {
      const bytes = new Uint8Array(data);
      const status = proto.Status.deserializeBinary(bytes)
      if(status.getCode() === proto.Status.Code.SUCCESS){
        if(hostState === HostState.CREATE_GAME){ // prevent race with sse
          localStorage.setItem('code', code)
          hostState = HostState.PROMPT
          startSSE()
        }
      } else {
        errMessage = status.getErrormessage() || 'Error (no message)'
      }
    });
  }

  // ====================== Prompts =====================

  let promptName

  const sendPrompt = () => {
    const msg = new proto.Prompt()

    msg.setName(promptName)
    msg.setCode(code)
    /* msg.setKey(key) */
    msg.setContestantsList(contestants)
    msg.setOptionsList(options)

    const body = msg.serializeBinary()

    fetch('/api/host/prompt', {
      method: 'POST',
      'content-type': 'application/x-protobuf',
      body,
    })
    .then(response => response.arrayBuffer())
    .then(data => {
      const bytes = new Uint8Array(data);
      const status = proto.Status.deserializeBinary(bytes)
      if(status.getCode() === proto.Status.Code.SUCCESS){
        hostState = HostState.RESOLVE
      } else {
        errMessage = status.getErrormessage() || 'Error (no message)'
      }
    });
  }

  // ==================== Contestants ===================

  $: contestants = JSON.parse(localStorage.getItem('contestants') || '[]')
  let addContestantName
  const addContestant = () => {
    contestants = [...contestants, addContestantName]
    prompt.contestants = contestants
    localStorage.setItem('contestants', JSON.stringify(contestants))
    addContestantName = ''
  }
  const removeContestant = contestantName => {

  }

  // ====================== Options =====================
  
  $: options = JSON.parse(localStorage.getItem('options') || '[]')
  let addOptionName
  const addOption = () => {
    options = [...options, addOptionName]
    prompt.options = options
    localStorage.setItem('options', JSON.stringify(options))
    addOptionName = ''
  }
  const removeOption = optionName => {

  }

  // ====================== RESOLVE =====================

  let playerScores = {}
  const setPlayerScore = (playerName, score) => {
    playerScores = {
      ...playerScores,
      [playerName]: score
    }
  }

  let playerAwardedPoints = {}
  const awardPlayerPoints = (playerName, points) => {
    if (!playerAwardedPoints.hasOwnProperty(playerName)) {
      playerAwardedPoints[playerName] = 0
    }

    playerAwardedPoints = {
      ...playerAwardedPoints,
      [playerName]: playerAwardedPoints[playerName] + points
    }
  }

  let hostAnswers = new Map()
  const chooseOption = (contestant, option) => {
    hostAnswers.set(contestant, option)
  }

  const resolve = () => {
    playerAwardedPoints = {}

    for (let [pName, pAnswers] of playerAnswers.entries()) {
      for (let pAnswer of pAnswers) {
        for (let [contestant, option] of hostAnswers.entries()) {
          if (
            pAnswer.contestant === contestant && 
            pAnswer.option === option
          ) {
            awardPlayerPoints(pName, 1)
          }
        }
      }
    }

    const resolveMsg = new proto.Resolve()
    resolveMsg.setCode(code)
    /* resolveMsg.setKey(key) */
    const playerScoresMsgs = []
  
    for (let playerName of playerNames) {
      const points = playerAwardedPoints[playerName] || 0

      setPlayerScore(playerName, playerScores[playerName] + points)

      const playerMsg = new proto.PlayerScore()
      playerMsg.setName(playerName)
      playerMsg.setScorechange(points)
      playerMsg.setNewscore(playerScores[playerName])
      playerScoresMsgs.push(playerMsg)
    }

    resolveMsg.setPlayerscoresList(playerScoresMsgs)

    const body = resolveMsg.serializeBinary()

    fetch('/api/host/resolve', {
      method: 'POST',
      'content-type': 'application/x-protobuf',
      body,
    })
    .then(response => response.arrayBuffer())
    .then(data => {
      const bytes = new Uint8Array(data);
      const status = proto.Status.deserializeBinary(bytes)
      if(status.getCode() === proto.Status.Code.SUCCESS){
        hostState = HostState.PROMPT
      } else {
        errMessage = status.getErrormessage() || 'Error (no message)'
      }
    })
  }


  // ===================== End Game =====================
  const endGame = () => {

    localStorage.clear()

  }

</script>

<main>
  {#if hostState == HostState.CREATE_GAME}
    <div class="create-game">
      <h2>Create Game</h2>
      <input type="text" placeholder="Code" bind:value={code} maxlength="6">
      <button class="big-btn" on:click={createGame}>Create</button>
    </div>
  {:else if hostState == HostState.PROMPT}
    <div class="host-prompt">
      <h3>Game Code: {code}</h3>

      <div class="prompt-form">
        <input type="text" placeholder="Prompt" bind:value={promptName}>

        <div class="input-contestants">
          <input type="text" placeholder="Add Contestant" bind:value={addContestantName}>
          <button class="big-btn" on:click={addContestant}>Add</button>
          <ul class="contestants">
            {#each contestants as contestant}
              <li>{contestant}
                <button class="tiny-btn" on:click={removeContestant(contestant)}>❌</button>
              </li>
            {/each}
          </ul>
        </div>

        <div class="input-options">
          <input type="text" placeholder="Add Option" bind:value={addOptionName}>
          <button class="big-btn" on:click={addOption}>Add</button>
          <ul class="options">
            {#each options as option}
              <li>{option}
                <button class="tiny-btn" on:click={removeOption(option)}>❌</button>
              </li>
            {/each}
          </ul>
        </div>

        <button class="big-btn" on:click={sendPrompt}>Send Prompt</button>
      </div>

      <button class="big-btn" on:click={endGame}>End Game</button>
    </div>
  {:else if hostState == HostState.RESOLVE}
    <div class="resolve">
      <h2>{promptName}</h2>
      <div class="contestants">
        {#each contestants as contestant}
          <div class="contestant">
            <h3>{contestant}</h3>
            <div class="options">
              {#each options as option}
                <div class="option">
                  <label>
                    <input type="radio" name="{contestant}" value="{option}" on:click={chooseOption(contestant,option)}> {option}
                  </label>
                </div>
              {/each}
            </div>
          </div>
        {/each}
      </div>

      <button on:click={resolve}>Resolve</button>
    </div>
  {/if}

  <p class="error">{errMessage}</p>

  <div class="player-name-list">
    <h3>Players</h3>
    <ul>
      {#each playerNames as playerName}
        <li class={ playersAnswered.includes(playerName) ? "answered" : ""}>
          <span class="player-name">
            {playerName} 
          </span>
          <span class="player-answered">
            {playersAnswered.includes(playerName) ? "✓" : "…"}
          </span>
          <span class="player-score">
            {playerScores[playerName]}
          </span>
          {#if playerAwardedPoints[playerName] > 0}
            <span class="player-score-increase">
              (⬆ {playerAwardedPoints[playerName]})
            </span>
          {/if}
        </li>
      {/each}
    </ul>
  </div>
</main>

<style>
  main {
    height: 100%;
    text-align: center;
    padding: 21px;
    margin: 0 auto;
    background-color: #333333;
    text-align: center;
    color: #ffffff;
    font-family: 'Space Mono', monospace;
  }
  /* h1 { */
  /*   font-size: 30px; */
  /* } */
  button.big-btn {
    margin: 0 auto;
    padding: 20px;
    color: #4F8132;
    display: block;
    width: 300px;
    font-size: 14px;
  }
  button.tiny-btn {
    color: #4F8132;
    font-size: 10px;
  }
  button:hover {
    text-decoration: none;
  }
  .error {
    color: #CC2010;
  }
</style>
