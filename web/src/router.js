import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

import Dashboard from './components/Dashboard';
import Playground from './components/Playground';

export default new VueRouter({
  routes: [
    { path: '/', redirect: '/dashboard' },
    { path: '/dashboard', component: Dashboard },
    { path: '/playground', component: Playground },
  ],
});
