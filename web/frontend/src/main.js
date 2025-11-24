import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.bundle.min.js'

import Login from './views/Login.vue'
import Customers from './views/Customers.vue'
import Accounts from './views/Accounts.vue'
import Analytics from './views/Analytics.vue'
import Dashboard from './views/Dashboard.vue'

const routes = [
  { path: '/', redirect: '/dashboard' },
  { path: '/login', component: Login },
  { path: '/dashboard', component: Dashboard, meta: { requiresAuth: true } },
  { path: '/customers', component: Customers, meta: { requiresAuth: true } },
  { path: '/accounts', component: Accounts, meta: { requiresAuth: true } },
  { path: '/analytics', component: Analytics, meta: { requiresAuth: true } }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Navigation guard for authentication
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  if (to.meta.requiresAuth && !token) {
    next('/login')
  } else {
    next()
  }
})

createApp(App).use(router).mount('#app')

