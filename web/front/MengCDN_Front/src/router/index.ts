import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import NFView from "@/views/NFView.vue";

const router = createRouter({

  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/gh/:owner/:repo@:branch',
      name: 'github',
      component: HomeView,
    },
    {
      path: '/gh/:owner/:repo@:branch/:pathMatch(.*)*',
      name: 'github404',
      component: NFView,
    },
    {
      path: '/gh/:owner/:repo/:pathMatch(.*)*',
      name: 'github404',
      component: NFView,
    },
    {
      path: '/gh/:owner/:pathMatch(.*)*',
      name: 'github404',
      component: NFView,
    },
    {
      path: '/gh/:pathMatch(.*)*',
      name: 'github404',
      component: NFView,
    }
  ]
})

export default router
