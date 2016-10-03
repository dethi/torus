import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

import Hello from './components/Hello';
import Playground from './components/Playground';

export default new VueRouter({
  routes: [
    { path: '/', component: Hello },
    { path: '/playground', component: Playground },
  ],
});
