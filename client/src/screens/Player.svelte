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
              comment: ''
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
            comment: s.comment
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
  .then(response => {
    if (!response.ok) throw response.statusText
    return response.arrayBuffer()
  })
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
    errMessage = ''
  })
  .catch(err => {
    console.error(err)
    errMessage = err.message || 'Error (no message)'
  })


  // ====================== PROMPT =====================

  let answers = {}
  const chooseOption = (contestant, option) => {
    answers = {
      ...answers,
      [contestant]: option,
    }
  }

  let comment = ''

  const submit = () => {
    const msg = new proto.PromptAnswer()
    const answerMsgs = []

    for (let [contestant, option] of Object.entries(answers)) {
      const answerMsg = new proto.PromptAnswer.Answer()
      answerMsg.setContestant(contestant)
      answerMsg.setOption(option)
      answerMsgs.push(answerMsg)
    }

    msg.setPlayername(name)
    msg.setCode(code)
    msg.setAnswersList(answerMsgs)
    msg.setComment(comment)

    const body = msg.serializeBinary()

    fetch('/api/player/answer', {
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
        gameState = GameState.RESOLVE
      } else {
        errMessage = status.getErrormessage() || 'Error (no message)'
      }
      answers = {}
      comment = ''
      errMessage = ''
    })
    .catch(err => {
      console.error(err)
      errMessage = err || 'Error (no message)'
    })
  }


  // ===================== Exit Game =====================
  const exitGame = () => {
    localStorage.clear()

    window.location.href = `/`
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
    <div class="status-box joining">Joining...</div>
  {:else if gameState == GameState.CONNECTING}
    <div class="status-box connecting">Connecting...</div>
  {:else if gameState == GameState.WAITING}
    <div class="status-box waiting">Waiting...</div>
  {:else if gameState == GameState.PROMPT}
    <div class="prompt">
      <h2>{prompt.name}</h2>
      <div class="contestants">
        {#each prompt.contestants as contestant}
          <div class="contestant">
            {#if contestant.thumbnail}
              <img class="contestant-thumbnail" src={contestant.thumbnail} alt={contestant.name} />
            {/if}
            <h3>{contestant.name}</h3>
            <div class="options">
              {#each prompt.options as option}
                <div 
                  class="option {answers[contestant.name] === option ? 'option-selected' : ''}"
                  on:click={chooseOption(contestant.name,option)}>
                    {answers[contestant.name] === option ? '✅' : ''}
                    {option}
                </div>
              {/each}
            </div>
          </div>
        {/each}
      </div>

      <h3>Comment</h3>
      <input class="comment-input" 
        placeholder={
          ["Extra predictions?","Say something nice","Just for fun"][parseInt(Math.random()*3)]
        } bind:value={comment}>

      <button on:click={submit}>Submit Answer</button>
    </div>
  {:else if gameState == GameState.RESOLVE}
    <div class="status-box resolve">Waiting for host to select winners...</div>
  {/if}

  <p class="error">{errMessage}</p>

  <h3>Scoreboard</h3>
  <div class="player-name-list">
    <ul>
      {#each Object.keys(playerScores) as playerName}
        <li class="scoreboard">
          <span class={playerName === name ? 'player-name current-player-name' : 'player-name'}>
            {playerName} 
          </span>
          <span class="player-score">
            {playerScores[playerName].newScore}
          </span>
          {#if playerScores[playerName].scoreChange > 0}
            <span class="player-score-increase">
              + {playerScores[playerName].scoreChange}
            </span>
          {/if}
          {#if playerScores[playerName].comment }
            <p class="player-comment">
              &ldquo;{playerScores[playerName].comment}&rdquo;
            </p>
          {/if}
        </li>
      {/each}
    </ul>
  </div>

  <div class="exit-game-container">
    <button class="exit-game-btn" on:click={exitGame}>Leave Game</button>
  </div>

</main>

<style>
  main {
    height: 100%;
    text-align: center;
    padding: 21px;
    margin: 0 auto;
    background-color: #fcfcfc;
    text-align: center;
    color: #232323;
    font-family: 'Pangolin', cursive;
  }
  h2 {
    font-family: 'Pangolin', cursive;
    font-size: 1.8em;
  }
  h3 {
    font-family: 'Pangolin', cursive;
    font-size: 3em;
    margin: 10px 0;
  }
  button {
    margin: 0 auto;
    margin-top: 40px;
    padding: 20px;
    display: block;
    width: 300px;
    font-family: 'Pangolin', cursive;
    font-weight: bold;
    font-size: 1.8em;
    background: rgba(66, 207, 127, 1);
    border: 0px solid;
    color: #fefefe;
  }
  button:hover {
    text-decoration: none;
  }
  .error {
    color: #CC2010;
  }
  .contestant {
    margin-bottom: 44px;
  }

  .status-box {
    width: 80%;
    height: 210px;
    background: #efefef;
    margin: auto;
    margin-top: 12px;
    margin-bottom: 12px;
    padding: 0 15px;
    display: flex;
    flex-flow: row nowrap;
    align-items: center;
    justify-content: center;
    font-size: 1.8em;
  }

  .my-score {
    font-size: 1.8em;
  }
  .my-score-status {
    font-size: 1.6em;
  }

  .comment-input {
    font-size: 1.8em;
    height: 120px;
  }

  .option {
    background: #fcfcfc;
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
    background: rgba(66, 207, 127, .2);
    border-color: rgba(66, 207, 127, 1);
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
  }
  .scoreboard:nth-of-type(2n+1) {
    background: #eee;
  }
  .player-name {
    flex: 0 0 50%;
    font-weight: bold;
    font-size: 2em;
  }
  .current-player-name {
    text-decoration: underline;
  }
  .player-score {
    flex: 0 0 20%;
    font-weight: bold;
    font-size: 2em;
  }
  .player-score-increase {
    flex: 0 0 30%;
    padding-top: 12px;
    color: #288a52;
  }
  .player-comment {
    flex: 0 0 100%;
  }

  button.exit-game-btn {
    margin: 0 auto;
    margin-top: 18px;
    padding: 20px;
    color: #fafafa;
    display: block;
    width: 120px;
    font-family: 'Pangolin', cursive;
    font-weight: bold;
    font-size: 1em;
    background: rgba(128, 23, 34, .8);
  }

  .options {
    display: flex;
    flex-flow: row wrap;
    justify-content: space-between;
  }
  .option {
    flex: 0 0 48%;
  }

  .contestant-thumbnail {
    max-width: 150px;
    display: block;
    margin: auto;
  }
</style>
