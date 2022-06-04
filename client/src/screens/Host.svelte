<script>
  const HostState = {
    CREATE_GAME: 0,
    PROMPT: 1,
    RESOLVE: 2,
  }

  let hostState = HostState.CREATE_GAME

  let code = localStorage.getItem('code')
  let key = localStorage.getItem('key')

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

  $: playerAnswers = {}
  $: playersAnswered = []
  let comments = {}
  const clearPlayerAnswers = () => {
    playerAnswers = {}
    playersAnswered = []
  }
  const setPlayerAnswer = ({playerName, answers, comment}) => {
    // save answer
    playerAnswers = {
      ...playerAnswers, 
      [playerName]: answers
    }

    // mark as answered
    playersAnswered = [...Object.keys(playerAnswers)]

    comments = {
      ...comments,
      [playerName]: comment
    }
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
    .then(response => {
      if (!response.ok) throw response.statusText
      return response.arrayBuffer()
    })
    .then(data => {
      const bytes = new Uint8Array(data);
      const hostCreateStatus = proto.HostCreateStatus.deserializeBinary(bytes)
      if(hostCreateStatus.getCode() === proto.HostCreateStatus.Code.SUCCESS){
        if(hostState === HostState.CREATE_GAME){ // prevent race with sse
          key = hostCreateStatus.getKey()
          localStorage.setItem('code', code)
          localStorage.setItem('key', key)
          hostState = HostState.PROMPT
          startSSE()
        }
      } else {
        errMessage = hostCreateStatus.getErrormessage() || 'Error (no message)'
      }
      comments = {}
      errMessage = ''
    })
    .catch(err => {
      console.error(err)
      errMessage = err || 'Error (no message)'
    })
  }

  const toUpperCase = e => {
    e.preventDefault();
    code = e.target.value = e.target.value.toUpperCase().replace(' ', '');
  }

  // ====================== Prompts =====================

  let promptName = localStorage.getItem('promptName') || ''

  const sendPrompt = () => {
    localStorage.setItem('promptName', promptName)

    const msg = new proto.Prompt()

    msg.setName(promptName)
    msg.setCode(code)
    msg.setKey(key)
    msg.setOptionsList(options)

    const contestantMsgs = contestants.map(({name, thumbnail}) => {
      const msg = new proto.Prompt.Contestant()
      msg.setName(name)
      msg.setThumbnail(thumbnail)
      return msg
    })
    msg.setContestantsList(contestantMsgs)

    const body = msg.serializeBinary()

    fetch('/api/host/prompt', {
      method: 'POST',
      'content-type': 'application/x-protobuf',
      body,
    })
    .then(response => {
      if (!response.ok) throw response.statusText
      return response.arrayBuffer()
    })
    .then(data => {
      const bytes = new Uint8Array(data);
      const status = proto.Status.deserializeBinary(bytes)
      clearPlayerAnswers()
      if(status.getCode() === proto.Status.Code.SUCCESS){
        hostState = HostState.RESOLVE
      } else {
        errMessage = status.getErrormessage() || 'Error (no message)'
      }
      errMessage = ''
    })
    .catch(err => {
      console.error(err)
      errMessage = err || 'Error (no message)'
    })
  }

  // ==================== Contestants ===================

  $: contestants = JSON.parse(localStorage.getItem('contestants') || '[]')
  let addContestantName
  let addContestantThumbnail
  const addContestant = () => {
    contestants = [...contestants, {
      name: addContestantName,
      thumbnail: addContestantThumbnail,
    }]
    prompt.contestants = contestants
    localStorage.setItem('contestants', JSON.stringify(contestants))
    addContestantName = ''
    addContestantThumbnail = ''
  }
  const removeContestant = contestantName => {
    contestants = [...contestants.filter(c => c.name !== contestantName)]
    prompt.contestants = contestants
    localStorage.setItem('contestants', JSON.stringify(contestants))
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
    options = [...options.filter(o => o !== optionName)]
    prompt.options = options
    localStorage.setItem('options', JSON.stringify(options))
  }

  // ====================== RESOLVE =====================

  let playerScores = {}
  const setPlayerScore = (playerName, score) => {
    playerScores = {
      ...playerScores,
      [playerName]: score
    }
  }

  const byScore = (playerName_a, playerName_b) => 
    playerScores[playerName_a] - playerScores[playerName_b]

  let playerAwardedPoints = {}
  const awardPlayerPoints = (playerName, points) => {
    if (!playerAwardedPoints.hasOwnProperty(playerName)) {
      playerAwardedPoints[playerName] = 0
    }

    playerAwardedPoints = {
      ...playerAwardedPoints,
      [playerName]: playerAwardedPoints[playerName] + points
    }

    playerNames = [...playerNamesSet].sort(byScore)
  }

  let hostAnswers = {}
  const chooseOption = (contestant, option) => {
    hostAnswers = {
      ...hostAnswers,
      [contestant]: option,
    }
  }

  const resolve = () => {
    playerAwardedPoints = {}

    for (let [pName, pAnswers] of Object.entries(playerAnswers)) {
      if (!pAnswers) { // user submitted empty answers
        pAnswers = []
      }
      for (let pAnswer of pAnswers) {
        for (let [contestant, option] of Object.entries(hostAnswers)) {
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
    resolveMsg.setKey(key)
    const playerScoresMsgs = []
  
    for (let playerName of playerNames) {
      const points = playerAwardedPoints[playerName] || 0

      setPlayerScore(playerName, playerScores[playerName] + points)

      const playerMsg = new proto.PlayerScore()
      playerMsg.setName(playerName)
      playerMsg.setScorechange(points)
      playerMsg.setNewscore(playerScores[playerName])
      playerMsg.setComment(comments[playerName])
      playerScoresMsgs.push(playerMsg)
    }

    resolveMsg.setPlayerscoresList(playerScoresMsgs)

    const body = resolveMsg.serializeBinary()

    fetch('/api/host/resolve', {
      method: 'POST',
      'content-type': 'application/x-protobuf',
      body,
    })
    .then(response => {
      if (!response.ok) throw response.statusText
      return response.arrayBuffer()
    })
    .then(data => {
      const bytes = new Uint8Array(data);
      const status = proto.Status.deserializeBinary(bytes)
      if(status.getCode() === proto.Status.Code.SUCCESS){
        hostState = HostState.PROMPT
      } else {
        errMessage = status.getErrormessage() || 'Error (no message)'
      }
      hostAnswers = {}
      errMessage = ''
    })
    .catch(err => {
      console.error(err)
      errMessage = err || 'Error (no message)'
    })
  }


  // ===================== End Game =====================
  const endGame = () => {
    localStorage.clear()

    window.location.href = `/host`
  }

</script>

<main>
  {#if hostState == HostState.CREATE_GAME}
    <div class="create-game">
      <h2>Create Game</h2>
      <input type="text" placeholder="Code" on:input={toUpperCase} bind:value={code} maxlength="6">
      <button class="big-btn" on:click={createGame}>Create</button>
    </div>
  {:else if hostState == HostState.PROMPT}
    <div class="host-prompt">
      <h3>Game Code: {code}</h3>

      <div class="prompt-section prompt-form">
        <label>Prompt<br>
          <input type="text" placeholder="Prompt" bind:value={promptName}>
        </label>
      </div>

      <div class="prompt-section input-contestants">
        <input type="text" placeholder="Add Contestant" bind:value={addContestantName}>
        <input type="text" placeholder="Thumbnail (optional)" bind:value={addContestantThumbnail}>
        <button class="medium-btn" on:click={addContestant}>Add</button>
        <ul class="contestants">
          {#each contestants as contestant}
            <li>
              {#if contestant.thumbnail}
                <img class="contestant-thumbnail" src={contestant.thumbnail} alt={contestant.name} />
              {/if}
              {contestant.name}
              <button class="tiny-btn" on:click={removeContestant(contestant.name)}>❌</button>
            </li>
          {/each}
        </ul>
      </div>

      <div class="prompt-section input-options">
        <input type="text" placeholder="Add Option" bind:value={addOptionName}>
        <button class="medium-btn" on:click={addOption}>Add</button>
        <ul class="options">
          {#each options as option}
            <li>{option}
              <button class="tiny-btn" on:click={removeOption(option)}>❌</button>
            </li>
          {/each}
        </ul>
      </div>

      <div class="prompt-section">
        <button class="big-btn" on:click={sendPrompt}>Send Prompt</button>
      </div>

      <div class="prompt-section ">
        <button class="end-game-btn" on:click={endGame}>End Game</button>
      </div>
    </div>
  {:else if hostState == HostState.RESOLVE}
    <div class="resolve">
      <h2>{promptName}</h2>
      <div class="contestants">
        {#each contestants as contestant}
          <div class="prompt-section contestant">
            {#if contestant.thumbnail}
              <img class="contestant-thumbnail" src={contestant.thumbnail} alt={contestant.name} />
            {/if}
            <h3>{contestant.name}</h3>
            <div class="options">
              {#each options as option}
                <div 
                  class="option {hostAnswers[contestant.name] === option ? 'option-selected' : ''}"
                  on:click={chooseOption(contestant.name,option)}>
                    {hostAnswers[contestant.name] === option ? '✅' : ''}
                    {option}
                </div>
              {/each}
            </div>
          </div>
        {/each}
      </div>

      <div class="prompt-section">
        <button class="big-btn" on:click={resolve}>Resolve</button>
      </div>
    </div>
  {/if}

  <p class="error">{errMessage}</p>

  {#if hostState != HostState.CREATE_GAME}
    <div class="player-name-list">
      <h3>Scoreboard</h3>
      <ul>
        {#each playerNames as playerName}
          <li class={ playersAnswered.includes(playerName) ? "scoreboard answered" : "scoreboard "}>
            <span class="player-name">
              {playerName} 
              <span class="player-answered">
                {playersAnswered.includes(playerName) ? "✓" : "…"}
              </span>
            </span>
            <span class="player-score">
              {playerScores[playerName]}
            </span>
            {#if playerAwardedPoints[playerName] > 0}
              <span class="player-score-increase">
                + {playerAwardedPoints[playerName]}
              </span>
            {/if}
            {#if comments[playerName] }
              <p class="player-comment">
                &ldquo;{comments[playerName]}&rdquo;
              </p>
            {/if}
          </li>
        {/each}
      </ul>
    </div>
  {/if}
</main>

<style>
  main {
    text-align: center;
    padding: 21px;
    margin: 0 auto;
    background-color: #333333;
    text-align: center;
    color: #ffffff;
    font-family: 'Pangolin', cursive;
  }
  h2 {
    font-family: 'Pangolin', cursive;
    font-size: 1.8em;
  }
  h3 {
    font-family: 'Pangolin', cursive;
    font-size: 3em;
  }
  button.big-btn {
    margin: 0 auto;
    padding: 20px;
    color: #4F8132;
    display: block;
    width: 300px;
    font-family: 'Pangolin', cursive;
    font-weight: bold;
    font-size: 1.8em;
    background: rgba(66, 127, 207, .4);
    border: 0px solid;
    color: #fefefe;
  }
  button.medium-btn {
    margin: 0 auto;
    padding: 14px;
    color: #4F8132;
    display: block;
    width: 200px;
    font-family: 'Pangolin', cursive;
    font-weight: bold;
    font-size: 1.6em;
    background: rgba(66, 127, 207, .4);
    border: 0px solid;
    color: #fefefe;
  }
  button.tiny-btn {
    color: #4F8132;
    font-size: 10px;
  }
  button.end-game-btn {
    margin: 0 auto;
    margin-top: 18px;
    padding: 20px;
    color: #fafafa;
    display: block;
    width: 140px;
    font-family: 'Pangolin', cursive;
    font-weight: bold;
    font-size: 1.2em;
    background: rgba(66, 127, 207, .4);
  }

  button:hover {
    text-decoration: none;
  }
  .error {
    color: #CC2010;
  }

  .prompt-section {
    border: 2px solid #444;
    padding: 18px;
    margin-bottom: 12px;
  }

  ul.contestants, ul.options {
    padding: 0;
    margin-bottom: 0;
    list-style: none;
  }

  .option {
    background: #101010;
    width: 80%;
    height: 60px;
    display: flex;
    flex-flow: row nowrap;
    align-items: center;
    justify-content: center;
    border: 4px solid #989898;
    margin: auto;
    margin-bottom: 12px;
    font-size: 1.8em;
  }
  .option-selected {
    background: rgba(66, 127, 207, .2);
    border-color: rgba(66, 127, 207, 1);
  }

  .player-name-list {
    text-align: left;
  }
  .player-name-list ul {
    padding: 0;
    margin: 0;
  }
  .scoreboard {
    list-style: none;
    padding: 12px;
    display: flex;
    flex-flow: row wrap;
    padding-left: 30px;
    background: #888;
  }
  .scoreboard:nth-of-type(2n+1) {
    background: #555;
  }
  .player-name {
    flex: 0 0 50%;
    font-weight: bold;
    font-size: 2em;
  }
  .player-score {
    flex: 0 0 20%;
    font-weight: bold;
    font-size: 2em;
  }
  .player-score-increase {
    flex: 0 0 30%;
    padding-top: 12px;
  }
  .player-comment {
    flex: 0 0 100%;
  }
  .contestant-thumbnail {
    max-width: 150px;
    display: block;
    margin: auto;
  }
</style>
