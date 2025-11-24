<template>
  <div class="container mt-5">
    <div class="row justify-content-center">
      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h3 class="text-center">Login</h3>
          </div>
          <div class="card-body">
            <form @submit.prevent="login">
              <div class="mb-3">
                <label for="username" class="form-label">Username</label>
                <input
                  type="text"
                  class="form-control"
                  id="username"
                  v-model="username"
                  required
                />
              </div>
              <div class="mb-3">
                <label for="password" class="form-label">Password</label>
                <input
                  type="password"
                  class="form-control"
                  id="password"
                  v-model="password"
                  required
                />
              </div>
              <div v-if="error" class="alert alert-danger">{{ error }}</div>
              <button type="submit" class="btn btn-primary w-100" :disabled="loading">
                {{ loading ? 'Logging in...' : 'Login' }}
              </button>
            </form>
            <div class="mt-3 text-center">
              <p>Don't have an account? <a href="#" @click.prevent="showRegister = true">Register</a></p>
            </div>
          </div>
        </div>

        <div class="card mt-3" v-if="showRegister">
          <div class="card-header">
            <h3 class="text-center">Register</h3>
          </div>
          <div class="card-body">
            <form @submit.prevent="register">
              <div class="mb-3">
                <label for="reg-username" class="form-label">Username</label>
                <input
                  type="text"
                  class="form-control"
                  id="reg-username"
                  v-model="regUsername"
                  required
                />
              </div>
              <div class="mb-3">
                <label for="reg-password" class="form-label">Password</label>
                <input
                  type="password"
                  class="form-control"
                  id="reg-password"
                  v-model="regPassword"
                  required
                  minlength="6"
                />
              </div>
              <div v-if="regError" class="alert alert-danger">{{ regError }}</div>
              <button type="submit" class="btn btn-success w-100" :disabled="regLoading">
                {{ regLoading ? 'Registering...' : 'Register' }}
              </button>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import apiClient from '../api/client'

export default {
  name: 'Login',
  setup() {
    const router = useRouter()
    const username = ref('')
    const password = ref('')
    const regUsername = ref('')
    const regPassword = ref('')
    const error = ref('')
    const regError = ref('')
    const loading = ref(false)
    const regLoading = ref(false)
    const showRegister = ref(false)

    const login = async () => {
      loading.value = true
      error.value = ''
      try {
        const response = await apiClient.post('/auth/login', {
          username: username.value,
          password: password.value
        })
        localStorage.setItem('token', response.data.token)
        router.push('/dashboard')
      } catch (err) {
        error.value = err.response?.data?.error || 'Login failed'
      } finally {
        loading.value = false
      }
    }

    const register = async () => {
      regLoading.value = true
      regError.value = ''
      try {
        await apiClient.post('/auth/register', {
          username: regUsername.value,
          password: regPassword.value
        })
        regError.value = ''
        showRegister.value = false
        alert('Registration successful! Please login.')
      } catch (err) {
        regError.value = err.response?.data?.error || 'Registration failed'
      } finally {
        regLoading.value = false
      }
    }

    return {
      username,
      password,
      regUsername,
      regPassword,
      error,
      regError,
      loading,
      regLoading,
      showRegister,
      login,
      register
    }
  }
}
</script>

