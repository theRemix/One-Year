<script>
  import "../pb/messages_pb.js"

  let code = ''
  let name = ''
  let errMessage = ''

  const join = () => {

    const msg = new proto.Join();

    msg.setPlayername(name);
    msg.setCode(code);

    const body = msg.serializeBinary();

    fetch('/api/join', {
      method: 'POST',
      'content-type': 'application/x-protobuf',
      body,
    })
    .then(response => response.arrayBuffer())
    .then(data => {
      const bytes = new Uint8Array(data);
      const status = proto.Status.deserializeBinary(bytes)
      if (status.getCode() === proto.Status.Code.SUCCESS) {
        localStorage.setItem('name', name)
        localStorage.setItem('code', code)
        window.location.href = `/player`
      } else {
        errMessage = status.getErrormessage() || 'Error (no message)'
      }
    });
  };
</script>

<main>
  <h1>Join Game</h1>

  <input type="text" placeholder="Code" bind:value={code}>
  <input type="text" placeholder="Your Name" bind:value={name}>

  <button on:click={join}>Join</button>

  <p class="error">{errMessage}</p>
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
