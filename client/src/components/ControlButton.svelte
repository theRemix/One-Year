<script>
  export let onClick;
  export let buttonText;
</script>

<button class="button green-button" on:click={onClick}>
  <!-- <slot></slot> -->
  {buttonText}
</button>