<script>
  import { Router, Link, Route } from "svelte-routing";
  import Root from "./screens/Root.svelte";
  import Join from "./screens/Join.svelte";
  import Player from "./screens/Player.svelte";
  import Host from "./screens/Host.svelte";
  export let url = "";
</script>

<Router {url}>
  <Route path="/" component={Root} />
  <Route path="/join" component={Join} />
  <Route path="/player" component={Player} />
  <Route path="/host" component={Host} />
</Router>
