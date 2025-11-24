<template>
  <div class="container mt-4">
    <div class="d-flex justify-content-between align-items-center mb-3">
      <h1>Customers</h1>
      <button class="btn btn-primary" @click="showCreateModal = true">Add Customer</button>
    </div>

    <table class="table table-striped">
      <thead>
        <tr>
          <th>ID</th>
          <th>Name</th>
          <th>Email</th>
          <th>Created At</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="customer in customers" :key="customer.id">
          <td>{{ customer.id }}</td>
          <td>{{ customer.name }}</td>
          <td>{{ customer.email }}</td>
          <td>{{ new Date(customer.created_at).toLocaleDateString() }}</td>
          <td>
            <button class="btn btn-sm btn-warning me-2" @click="editCustomer(customer)">Edit</button>
            <button class="btn btn-sm btn-danger" @click="deleteCustomer(customer.id)">Delete</button>
          </td>
        </tr>
      </tbody>
    </table>

    <!-- Create/Edit Modal -->
    <div class="modal" :class="{ show: showCreateModal || showEditModal, 'd-block': showCreateModal || showEditModal }" tabindex="-1">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">{{ showEditModal ? 'Edit Customer' : 'Create Customer' }}</h5>
            <button type="button" class="btn-close" @click="closeModal"></button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="saveCustomer">
              <div class="mb-3">
                <label class="form-label">Name</label>
                <input type="text" class="form-control" v-model="form.name" required />
              </div>
              <div class="mb-3">
                <label class="form-label">Email</label>
                <input type="email" class="form-control" v-model="form.email" required />
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
  name: 'Customers',
  setup() {
    const customers = ref([])
    const showCreateModal = ref(false)
    const showEditModal = ref(false)
    const form = ref({ name: '', email: '' })
    const editingId = ref(null)
    const error = ref('')

    const loadCustomers = async () => {
      try {
        const response = await apiClient.get('/customers')
        customers.value = response.data
      } catch (err) {
        error.value = 'Failed to load customers'
      }
    }

    const saveCustomer = async () => {
      error.value = ''
      try {
        if (editingId.value) {
          await apiClient.put(`/customers/${editingId.value}`, form.value)
        } else {
          await apiClient.post('/customers', form.value)
        }
        closeModal()
        loadCustomers()
      } catch (err) {
        error.value = err.response?.data?.error || 'Failed to save customer'
      }
    }

    const editCustomer = (customer) => {
      editingId.value = customer.id
      form.value = { name: customer.name, email: customer.email }
      showEditModal.value = true
    }

    const deleteCustomer = async (id) => {
      if (!confirm('Are you sure you want to delete this customer?')) return
      try {
        await apiClient.delete(`/customers/${id}`)
        loadCustomers()
      } catch (err) {
        error.value = err.response?.data?.error || 'Failed to delete customer'
      }
    }

    const closeModal = () => {
      showCreateModal.value = false
      showEditModal.value = false
      form.value = { name: '', email: '' }
      editingId.value = null
      error.value = ''
    }

    onMounted(() => {
      loadCustomers()
    })

    return {
      customers,
      showCreateModal,
      showEditModal,
      form,
      error,
      saveCustomer,
      editCustomer,
      deleteCustomer,
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

