<template>
  <body id="app">
    <AppHeader/>
    <router-view/>
    <AppFooter :version="version"/>
  </body>
</template>

<script>
import axios from 'axios';
import AppHeader from './components/Header';
import AppFooter from './components/Footer';

export default {
  components: {
    AppHeader,
    AppFooter,
  },
  data() {
    return {
      version: '',
    };
  },
  methods: {
    fetchVersion() {
      const vm = this;
      axios.get('/api/version').then(res => {
        vm.version = res.data.revision;
      });
    },
  },
  created() {
    this.fetchVersion();
  },
};
</script>

<style>
</style>
