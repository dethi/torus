<template>
  <section class="section">
    <div class="columns">
      <div class="column is-half">
        <nav class="level">
          <div class="level-item has-text-centered">
            <p class="heading">Torrents</p>
            <p class="title">N/A</p>
          </div>
          <div class="level-item has-text-centered">
            <p class="heading">Disk Space</p>
            <p class="title">N/A</p>
          </div>
        </nav>
      </div>

      <div class="column is-half">
        <article class="message is-warning">
          <div class="message-header">Warning</div>
          <div class="message-body">
            The torrents are deleted after 48 hours due to server migration.
          </div>
        </article>
      </div>
    </div>

    <Torrent>
  </section>
</template>

<script>
import Torrent from './Torrent';

export default {
  components: {
    Torrent,
  },
};
</script>

<style scoped>
</style>
