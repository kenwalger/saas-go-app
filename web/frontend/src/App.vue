<template>
  <div id="app">
    <nav class="navbar navbar-expand-lg navbar-dark bg-primary" v-if="isAuthenticated">
      <div class="container-fluid">
        <a class="navbar-brand" href="#">SaaS Go App</a>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarNav">
          <ul class="navbar-nav">
            <li class="nav-item">
              <router-link to="/dashboard" class="nav-link">Dashboard</router-link>
            </li>
            <li class="nav-item">
              <router-link to="/customers" class="nav-link">Customers</router-link>
            </li>
            <li class="nav-item">
              <router-link to="/accounts" class="nav-link">Accounts</router-link>
            </li>
            <li class="nav-item">
              <router-link to="/analytics" class="nav-link">Analytics</router-link>
            </li>
          </ul>
          <ul class="navbar-nav ms-auto">
            <li class="nav-item">
              <button class="btn btn-outline-light btn-sm" @click="logout">Logout</button>
            </li>
          </ul>
        </div>
      </div>
    </nav>
    <router-view />
  </div>
</template>

<script>
import { computed } from 'vue'
import { useRouter } from 'vue-router'

export default {
  name: 'App',
  setup() {
    const router = useRouter()
    const isAuthenticated = computed(() => {
      return !!localStorage.getItem('token')
    })

    const logout = () => {
      localStorage.removeItem('token')
      router.push('/login')
    }

    return {
      isAuthenticated,
      logout
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
</style>

