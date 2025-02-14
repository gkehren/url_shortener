<script>
  export let shortUrl;
</script>

{#if shortUrl}
  <div class="mt-4">
    Short URL: <a href={`http://localhost:8080/${shortUrl}`} class="text-blue-500">{shortUrl}</a>
  </div>
{/if}
