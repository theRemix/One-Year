<script>
  const join = () => {
    window.location.href = "/join";
  };
</script>

<main>
  <div class="content">
    <h1>One Year</h1>

    <div class="header-img">
      <img src="/assets/header.jpg" alt="header image" />
    </div>

    <button on:click={join}>Join Game</button>
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
  img {
      width: 35%;
  }
  h1 {
    font-family: "Press Start 2P";
  }
  :global(p) {
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
