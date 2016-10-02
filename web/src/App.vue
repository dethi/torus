<template>
  <body id="app">
    <AppHeader/>
    <router-view/>
    <AppFooter version="0.0.1"/>
  </body>
</template>

<script>
import AppHeader from './components/Header';
import AppFooter from './components/Footer';

export default {
  components: {
    AppHeader,
    AppFooter,
  },
};
</script>

<style>
</style>
