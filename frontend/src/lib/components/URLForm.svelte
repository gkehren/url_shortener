<script>
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  let longUrl = '';

  async function handleSubmit(event) {
    event.preventDefault();
    const response = await fetch('http://localhost:8080/shorten', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ long_url: longUrl }),
    });
    const data = await response.json();
    dispatch('urlShortened', data.short_url);
  }
</script>

<form on:submit|preventDefault={handleSubmit} class="space-y-4">
  <input
    type="text"
    bind:value={longUrl}
    placeholder="Enter long URL"
    class="border p-2 rounded w-full"
  />
  <button type="submit" class="bg-blue-500 text-white p-2 rounded">Shorten</button>
</form>
