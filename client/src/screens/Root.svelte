<script>
  const join = () => {
    window.location.href = "/join";
  };
</script>

<main>
  <div class="content">
    <h1>One Year</h1>
    <button on:click={join}>Join Game</button>
  </div>
</main>

<style>
  h1 {
    font-family: "Press Start 2P";
  }
  :global(p) {
    font-family: "Space Mono", monospace;
  }
  main {
    text-align: center;
    margin: 0 auto;
    text-align: center;
    color: #ffffff;
    font-family: "Space Mono", monospace;
  }
  button {
    margin-top: 50px;
    width: 249px;
    height: 43px;
    font-size: 10px;
    cursor: pointer;
  }
</style>
