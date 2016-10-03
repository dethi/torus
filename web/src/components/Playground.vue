<template>
  <div class="playground">
    <p>
      Search:
      <input v-model="query">
    </p>

    <p>{{ msg }}</p>

    <div>
      <p v-for="hit in hits">{{ hit.name }} - {{ hit.size }}</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import _ from 'lodash';

export default {
  data() {
    return {
      query: '',
      msg: '',
      hits: [],
    };
  },
  watch: {
    query() {
      this.getHits();
    },
  },
  methods: {
    getHits: _.debounce(function search() {
      const vm = this;
      const q = vm.query.trim();
      if (q === '') {
        return;
      }

      vm.msg = 'Searching...';

      axios.get('/api/search/cpasbien', {
        params: { q },
      }).then(res => {
        vm.hits = res.data.hits;
        vm.msg = '';
      });
    }, 500),
  },
};
</script>

<style scoped>
</style>
