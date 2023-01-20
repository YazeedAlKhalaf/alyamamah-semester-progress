<script lang="ts">
  export let text: string;
  export let isSelected: boolean;
  export let onClick: () => void;
</script>

<button class={`chip ${isSelected ? "chip-selected" : ""}`} on:click={onClick}>
  {text}
</button>

<style>
  .chip {
    display: inline-block;
    padding: 0.25rem 0.5rem;
    margin: 0.25rem;
    border-radius: 0.25rem;
    background-color: #eaeaea;
    color: #0f0f0f;
    border: none;
    font-size: 1rem;
  }

  .chip:hover {
    background-color: #dcdcdc;
    color: #0f0f0f;
    cursor: pointer;
  }

  .chip-selected {
    background-color: #0f0f0f;
    color: #ffffff;
  }
</style>
