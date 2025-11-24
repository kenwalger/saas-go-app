<template>
  <div class="container mt-4">
    <h1>Analytics</h1>
    <div class="row mt-4">
      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h5>Overall Statistics</h5>
          </div>
          <div class="card-body">
            <table class="table">
              <tr>
                <td><strong>Total Customers:</strong></td>
                <td>{{ analytics.total_customers || 0 }}</td>
              </tr>
              <tr>
                <td><strong>Total Accounts:</strong></td>
                <td>{{ analytics.total_accounts || 0 }}</td>
              </tr>
              <tr>
                <td><strong>Active Accounts:</strong></td>
                <td>{{ analytics.active_accounts || 0 }}</td>
              </tr>
              <tr>
                <td><strong>Inactive Accounts:</strong></td>
                <td>{{ analytics.inactive_accounts || 0 }}</td>
              </tr>
              <tr>
                <td><strong>Avg Accounts per Customer:</strong></td>
                <td>{{ analytics.avg_accounts_per_customer?.toFixed(2) || 0 }}</td>
              </tr>
            </table>
          </div>
        </div>
      </div>
      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h5>Customer Analytics</h5>
          </div>
          <div class="card-body">
            <div class="mb-3">
              <label class="form-label">Customer ID</label>
              <input
                type="number"
                class="form-control"
                v-model="customerId"
                @input="loadCustomerAnalytics"
                placeholder="Enter customer ID"
              />
            </div>
            <div v-if="customerAnalytics">
              <table class="table">
                <tr>
                  <td><strong>Total Accounts:</strong></td>
                  <td>{{ customerAnalytics.total_accounts || 0 }}</td>
                </tr>
                <tr>
                  <td><strong>Active Accounts:</strong></td>
                  <td>{{ customerAnalytics.active_accounts || 0 }}</td>
                </tr>
                <tr>
                  <td><strong>Inactive Accounts:</strong></td>
                  <td>{{ customerAnalytics.inactive_accounts || 0 }}</td>
                </tr>
              </table>
            </div>
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
  name: 'Analytics',
  setup() {
    const analytics = ref({})
    const customerAnalytics = ref(null)
    const customerId = ref('')

    const loadAnalytics = async () => {
      try {
        const response = await apiClient.get('/analytics')
        analytics.value = response.data
      } catch (error) {
        console.error('Failed to load analytics:', error)
      }
    }

    const loadCustomerAnalytics = async () => {
      if (!customerId.value) {
        customerAnalytics.value = null
        return
      }
      try {
        const response = await apiClient.get(`/analytics/customers/${customerId.value}`)
        customerAnalytics.value = response.data
      } catch (error) {
        console.error('Failed to load customer analytics:', error)
        customerAnalytics.value = null
      }
    }

    onMounted(() => {
      loadAnalytics()
    })

    return {
      analytics,
      customerAnalytics,
      customerId,
      loadCustomerAnalytics
    }
  }
}
</script>

