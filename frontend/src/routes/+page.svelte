<script>
  import URLForm from '$lib/components/URLForm.svelte';
  import URLResult from '$lib/components/URLResult.svelte';

  let shortUrl = '';

  function handleUrlShortened(event) {
    shortUrl = event.detail;
  }
</script>

<main class="p-4">
  <h1 class="text-2xl font-bold mb-4">URL Shortener</h1>
  <URLForm on:urlShortened={handleUrlShortened} />
  <URLResult {shortUrl} />
</main>
