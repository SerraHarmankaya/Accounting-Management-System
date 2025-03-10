<template>
  <q-page padding>
    <q-table
      flat
      bordered
      title="Müşteri Listesi"
      :rows="clients"
      :columns="columns"
      row-key="id"
    />
  </q-page>
</template>

<script>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useQuasar } from 'quasar'

export default {
  setup() {
    const $q = useQuasar()
    const clients = ref([])

    // API'den veri çekme fonksiyonu
    const fetchClients = async () => {
      try {
        const response = await axios.get('http://localhost:9000/clients', {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`,
          },
        })
        clients.value = response.data
      } catch (error) {
        $q.notify({
          type: 'negative',
          message: 'Veriler alınırken hata oluştu!',
        })
        return error
      }
    }

    onMounted(fetchClients)

    // Tablo Sütunları
    const columns = [
      { name: 'id', label: 'ID', align: 'left', field: (row) => row.id, sortable: true },
      { name: 'name', label: 'Adı', align: 'left', field: 'name', sortable: true },
      { name: 'email', label: 'E-Posta', align: 'left', field: 'email' },
      { name: 'role', label: 'Rol', align: 'center', field: 'role', sortable: true },
    ]

    return { clients, columns }
  },
}
</script>
