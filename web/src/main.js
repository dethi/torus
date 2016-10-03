import Vue from 'vue';
import App from './App';
import router from './router';
import axios from 'axios';

// In development, I use `npm run dev` for live reloading, and a instance
// of `torus` for the API.
if (process.env.NODE_ENV === 'development') {
  axios.defaults.baseURL = 'http://localhost:8000';
}

/* eslint-disable no-new */
const app = new Vue({
  router,
  ...App,
});

app.$mount('#app');
