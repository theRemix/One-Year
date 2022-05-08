<script>
  const name = localStorage.getItem('name')
  const code = localStorage.getItem('code')

  const GameState = {
    JOINING: 0,
    CONNECTING: 1,
    WAITING: 2,
    PROMPT: 3,
    RESOLVE: 4,
  }

  $: gameState = GameState.JOINING

  $: errMessage = ""

  // ==================== State =====================

  const playerNamesSet = new Set()
  let playerNames = []
  let playerScores = {}
  let curPlayerScore = 0
  let curPlayerStatus = ''

  const addPlayerName = playerName => {
    playerScores = {
      ...playerScores,
      [playerName]: {
        newScore: 0,
        scoreChange: 0,
      },
    }

    playerNamesSet.add(playerName)
    playerNames = [...playerNamesSet]
  }

  let prompt = {
    contestants: [],
    options: [],
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
        addPlayerName(payload.playerName)
        break
      case proto.PlayerEvent.Op.PROMPT:
        gameState = GameState.PROMPT
        prompt = payload.prompt

        // reset some props from last round
        curPlayerStatus = ''
        playerScores = Object.entries(playerScores)
          .reduce((ps,[name, score]) => ({
            ...ps,
            [name]: {
              ...score,
              scoreChange: 0,
            }
          }), {})
        break
      case proto.PlayerEvent.Op.RESOLVED:
        gameState = GameState.WAITING
        playerScores = payload.playerScores.reduce((ps,s) => ({
          ...ps,
          [s.name]: { // 0 is omitted
            scoreChange: s.scoreChange || 0,
            newScore: s.newScore || 0, 
          }
        }), {})
        curPlayerScore = playerScores[name].newScore
        curPlayerStatus = `You earned ${playerScores[name].scoreChange} points`
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

  msg.setPlayername(name);
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


  // ====================== PROMPT =====================

  let answers = new Map()
  const chooseOption = (contestant, option) => {
    answers.set(contestant, option)
  }

  const submit = () => {
    const msg = new proto.PromptAnswer()
    const answerMsgs = []

    for (let [contestant, option] of answers.entries()) {
      const answerMsg = new proto.PromptAnswer.Answer()
      answerMsg.setContestant(contestant)
      answerMsg.setOption(option)
      answerMsgs.push(answerMsg)
    }

    msg.setPlayername(name)
    msg.setCode(code)
    msg.setAnswersList(answerMsgs)

    const body = msg.serializeBinary()

    fetch('/api/player/answer', {
      method: 'POST',
      'content-type': 'application/x-protobuf',
      body,
    })
    .then(response => response.arrayBuffer())
    .then(data => {
      const bytes = new Uint8Array(data);
      const status = proto.Status.deserializeBinary(bytes)
      if(status.getCode() === proto.Status.Code.SUCCESS){
        gameState = GameState.RESOLVE
      } else {
        errMessage = status.getErrormessage() || 'Error (no message)'
      }
    })
  }



</script>

<main>
  {#if [GameState.WAITING, GameState.PROMPT, GameState.RESOLVE].includes(gameState) }
    <div class="my-scores">
      <div class="my-score-status">{curPlayerStatus}</div>
      <div class="my-score">Score: {curPlayerScore}</div>
    </div>
  {/if}

  {#if gameState == GameState.JOINING}
    <div class="joining">Joining...</div>
  {:else if gameState == GameState.CONNECTING}
    <div class="connecting">Connecting...</div>
  {:else if gameState == GameState.WAITING}
    <div class="waiting">Waiting...</div>
  {:else if gameState == GameState.PROMPT}
    <div class="prompt">
      <h2>{prompt.name}</h2>
      <div class="contestants">
        {#each prompt.contestants as contestant}
          <div class="contestant">
            <h3>{contestant}</h3>
            <div class="options">
              {#each prompt.options as option}
                <div class="option">
                  <label>
                    <input type="radio" name={contestant} value={option} on:click={chooseOption(contestant,option)}> {option}
                  </label>
                </div>
              {/each}
            </div>
          </div>
        {/each}
      </div>

      <button on:click={submit}>Submit Answer</button>
    </div>
  {:else if gameState == GameState.RESOLVE}
    <div class="resolve">Waiting for host to select winners...</div>
  {/if}

  <p class="error">{errMessage}</p>

  <div class="player-name-list">
    <h3>Players</h3>
    <ul>
      {#each playerNames as playerName}
        <li>
          <span class="player-name">
            {playerName} 
          </span>
          <span class="player-score">
            {playerScores[playerName].newScore}
          </span>
          {#if playerScores[playerName].scoreChange > 0}
            <span class="player-score-increase">
              (⬆ {playerScores[playerName].scoreChange})
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
