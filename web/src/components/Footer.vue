<template>
  <footer class="footer">
    <div class="container">
      <div class="content has-text-centered">
        <p><em>{{ version }}</em></p>
        <p>
          <a class="icon" href="https://github.com/dethi/torus">
            <i class="fa fa-github"></i>
          </a>
        </p>
      </div>
    </div>
  </footer>
</template>

<script>
export default {
  props: {
    version: {
      type: String,
      required: true,
    },
  },
};
</script>

<style scoped>
</style>
