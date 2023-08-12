<script lang="ts">
  export let title: string;
  export let body: string;
  export let isHighlighted: boolean;
</script>

<li class="link-card" style={isHighlighted ? "background-position: 0;" : ""}>
  <div>
    <h2 style={isHighlighted ? "color: #f8933e;" : ""}>
      {title}
    </h2>
    <p>
      {body}
    </p>
  </div>
</li>

<style>
  :root {
    --link-gradient: linear-gradient(
      45deg,
      #f8933e,
      #feb351 30%,
      var(--color-border) 60%
    );
  }

  .link-card {
    list-style: none;
    display: flex;
    padding: 0.15rem;
    background-image: var(--link-gradient);
    background-size: 400%;
    border-radius: 0.5rem;
    background-position: 100%;
    transition: background-position 0.6s cubic-bezier(0.22, 1, 0.36, 1);
  }

  .link-card > div {
    width: 100%;
    text-decoration: none;
    line-height: 1.4;
    padding: 1em 1.3em;
    border-radius: 0.35rem;
    color: var(--text-color);
    background-color: white;
    opacity: 0.8;
  }

  h2 {
    margin: 0;
    transition: color 0.6s cubic-bezier(0.22, 1, 0.36, 1);
  }

  p {
    margin-top: 0.75rem;
    margin-bottom: 0;
  }
</style>
