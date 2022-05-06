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
  }

  const playersAnsweredSet = new Set()
  $: playersAnswered = []
  const setPlayerAnswered = playerName => {
    playersAnsweredSet.add(playerName)
    playersAnswered = [...playersAnsweredSet]
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
            gameState = GameState.PROMPT
          } 
          addPlayerName(payload.playerName)
          break
        case proto.PlayerEvent.Op.ANSWERED:
          setPlayerAnswered(payload.playerName)
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

    msg.setPlayername(promptName)
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
    localStorage.setItem('options', JSON.stringify(options))
    addOptionName = ''
  }
  const removeOption = optionName => {

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
  {/if}

  <p class="error">{errMessage}</p>

  <div class="player-name-list">
    <h3>Connected Players</h3>
    <ul>
      {#each playerNames as playerName}
        <li class={ playersAnswered.includes(playerName) ? "answered" : ""}>{playerName} { playersAnswered.includes(playerName) ? "✓" : "…"}</li>
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
