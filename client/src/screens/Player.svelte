<script>
  const params = new URLSearchParams(window.location.search)
  const name = localStorage.getItem('name')
  const code = localStorage.getItem('code')

  const GameState = {
    JOINING: 0,
    CONNECTING: 1,
    WAITING: 2,
  }

  $: gameState = GameState.JOINING

  $: errMessage = ""

  // ==================== State =====================

  const playerNamesSet = new Set()
  $: playerNames = []
  const addPlayerName = playerName => {
    playerNamesSet.add(playerName)
    playerNames = [...playerNamesSet]
  }

  // ====================== SSE =====================
  
  const events = new EventSource(`/events/player/${code}`)

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


  // ====================== JOIN =====================

  const msg = new proto.Join();

  msg.setName(name);
  msg.setCode(code);

  const body = msg.serializeBinary()

  fetch('/api/player/ready', {
    method: 'POST',
    'content-type': 'application/x-protobuf',
    body,
  })
  .then(response => response.arrayBuffer())
  .then(data => {
    const bytes = new Uint8Array(data);
    const status = proto.Status.deserializeBinary(bytes)
    if(status.getCode() === proto.Status.Code.SUCCESS){
      if(gameState === GameState.JOINING){ // prevent race with sse
        gameState = GameState.CONNECTING
      }
    } else {
      errMessage = status.getErrormessage() || 'Error (no message)'
    }
  });



</script>

<main>
  {#if gameState == GameState.JOINING}
    <div class="joining">Joining...</div>
  {:else if gameState == GameState.CONNECTING}
    <div class="connecting">Connecting...</div>
  {:else if gameState == GameState.WAITING}
    <div class="waiting">Waiting...</div>
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
  h1 {
    font-size: 30px;
  }
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
