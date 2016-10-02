import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

import Hello from './components/Hello';

export default new VueRouter({
  routes: [
    { path: '/', component: Hello },
  ],
});
