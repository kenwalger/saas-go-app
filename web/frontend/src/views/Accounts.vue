<template>
  <div class="container mt-4">
    <div class="d-flex justify-content-between align-items-center mb-3">
      <h1>Accounts</h1>
      <button class="btn btn-primary" @click="showCreateModal = true">Add Account</button>
    </div>

    <table class="table table-striped">
      <thead>
        <tr>
          <th>ID</th>
          <th>Customer ID</th>
          <th>Name</th>
          <th>Status</th>
          <th>Created At</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="account in accounts" :key="account.id">
          <td>{{ account.id }}</td>
          <td>{{ account.customer_id }}</td>
          <td>{{ account.name }}</td>
          <td>
            <span class="badge" :class="account.status === 'active' ? 'bg-success' : 'bg-secondary'">
              {{ account.status }}
            </span>
          </td>
          <td>{{ new Date(account.created_at).toLocaleDateString() }}</td>
          <td>
            <button class="btn btn-sm btn-warning me-2" @click="editAccount(account)">Edit</button>
            <button class="btn btn-sm btn-danger" @click="deleteAccount(account.id)">Delete</button>
          </td>
        </tr>
      </tbody>
    </table>

    <!-- Create/Edit Modal -->
    <div class="modal" :class="{ show: showCreateModal || showEditModal, 'd-block': showCreateModal || showEditModal }" tabindex="-1">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">{{ showEditModal ? 'Edit Account' : 'Create Account' }}</h5>
            <button type="button" class="btn-close" @click="closeModal"></button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="saveAccount">
              <div class="mb-3">
                <label class="form-label">Customer ID</label>
                <input type="number" class="form-control" v-model="form.customer_id" required :disabled="showEditModal" />
              </div>
              <div class="mb-3">
                <label class="form-label">Name</label>
                <input type="text" class="form-control" v-model="form.name" required />
              </div>
              <div class="mb-3">
                <label class="form-label">Status</label>
                <select class="form-select" v-model="form.status" required>
                  <option value="active">Active</option>
                  <option value="inactive">Inactive</option>
                </select>
              </div>
              <div v-if="error" class="alert alert-danger">{{ error }}</div>
              <button type="submit" class="btn btn-primary">Save</button>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import apiClient from '../api/client'

export default {
  name: 'Accounts',
  setup() {
    const accounts = ref([])
    const showCreateModal = ref(false)
    const showEditModal = ref(false)
    const form = ref({ customer_id: '', name: '', status: 'active' })
    const editingId = ref(null)
    const error = ref('')

    const loadAccounts = async () => {
      try {
        const response = await apiClient.get('/accounts')
        accounts.value = response.data
      } catch (err) {
        error.value = 'Failed to load accounts'
      }
    }

    const saveAccount = async () => {
      error.value = ''
      try {
        if (editingId.value) {
          await apiClient.put(`/accounts/${editingId.value}`, {
            name: form.value.name,
            status: form.value.status
          })
        } else {
          await apiClient.post('/accounts', {
            customer_id: parseInt(form.value.customer_id),
            name: form.value.name,
            status: form.value.status
          })
        }
        closeModal()
        loadAccounts()
      } catch (err) {
        error.value = err.response?.data?.error || 'Failed to save account'
      }
    }

    const editAccount = (account) => {
      editingId.value = account.id
      form.value = {
        customer_id: account.customer_id,
        name: account.name,
        status: account.status
      }
      showEditModal.value = true
    }

    const deleteAccount = async (id) => {
      if (!confirm('Are you sure you want to delete this account?')) return
      try {
        await apiClient.delete(`/accounts/${id}`)
        loadAccounts()
      } catch (err) {
        error.value = err.response?.data?.error || 'Failed to delete account'
      }
    }

    const closeModal = () => {
      showCreateModal.value = false
      showEditModal.value = false
      form.value = { customer_id: '', name: '', status: 'active' }
      editingId.value = null
      error.value = ''
    }

    onMounted(() => {
      loadAccounts()
    })

    return {
      accounts,
      showCreateModal,
      showEditModal,
      form,
      error,
      saveAccount,
      editAccount,
      deleteAccount,
      closeModal
    }
  }
}
</script>

<style scoped>
.modal.show {
  background-color: rgba(0, 0, 0, 0.5);
}
</style>

