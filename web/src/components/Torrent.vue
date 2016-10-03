<template>
  <div class="container">
    <table class="table is-striped">
      <thead>
        <tr>
          <th>Name</th>
          <th>Since</th>
          <th>Download</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in records">
          <td>{{ item.torrents[0].name }}</td>
          <td>{{ item.request.request_time }}</td>
          <td class="is-icon">
            <a :href="downloadUrl(item)">
              <i class="fa fa-globe"></i>
            </a>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      records: [],
    };
  },
  methods: {
    downloadUrl(item) {
      return `/api/file/${item.torrents[0].info_hash}`;
    },
    getRecords() {
      const vm = this;
      axios.get('/api/torrent').then(res => {
        vm.records = res.data.data;
      });
    },
  },
  mounted() {
    this.getRecords();
  },
};
</script>

<style scoped>
</style>
