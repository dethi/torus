import Vue from 'vue';
import App from './App';
import router from './router';

/* eslint-disable no-new */
const app = new Vue({
  router,
  ...App,
});

app.$mount('#app');
