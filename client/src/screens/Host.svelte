<script>
  const HostState = {
    CREATE: 0,
    WAITING: 1,
  }

  let code = localStorage.getItem('code')
  let hostState = HostState.CREATE
  let errMessage = ""

  /* if (code) { // restore previous session */
  /*   hostState = HostState.WAITING */
  /*   startSSE() */
  /* } */

  // ==================== State =====================

  const playerNamesSet = new Set()
  $: playerNames = []
  const addPlayerName = playerName => {
    playerNamesSet.add(playerName)
    playerNames = [...playerNamesSet]
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
            gameState = GameState.WAITING
          } 
          addPlayerName(payload.name)
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
        if(hostState === HostState.CREATE){ // prevent race with sse
          hostState = HostState.WAITING
          startSSE()
        }
      } else {
        errMessage = status.getErrormessage() || 'Error (no message)'
      }
    });
  }



</script>

<main>
  {#if hostState == HostState.CREATE}
    <div class="create-game">
      <h2>Create Game</h2>
      <input type="text" placeholder="Code" bind:value={code} maxlength="6">
      <button on:click={createGame}>Create</button>
    </div>
  {:else if hostState == HostState.WAITING}
    <div class="host-waiting">
      <h3>Game Code: {code}</h3>
    </div>
  {/if}

  <p class="error">{errMessage}</p>

  <div class="player-name-list">
    <h3>Connected Players</h3>
    <ul>
      {#each playerNames as playerName}
        <li>{playerName}</li>
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
  button {
    margin: 0 auto;
    margin-top: 100px;
    padding: 20px;
    color: #4F8132;
    display: block;
    width: 300px;
    font-size: 14px;
  }
  button:hover {
    text-decoration: none;
  }
  .error {
    color: #CC2010;
  }
</style>
