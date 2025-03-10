<template>
  <q-page padding>
    <q-table
      flat
      bordered
      title="Müşteri Listesi"
      :rows="clients"
      :columns="columns"
      row-key="id"
    >
      <template #body-cell-actions="props">
        <q-td :props="props">
          <q-btn
            icon="delete"
            color="negative"
            round
            dense
            flat
            @click="confirmDelete(props.row.id)"
          />
        </q-td>
      </template>
    </q-table>

    <!-- Onay Diyaloğu -->
    <q-dialog v-model="dialogVisible">
      <q-card>
        <q-card-section class="row items-center">
          <q-avatar icon="warning" color="orange" text-color="white" />
          <span class="q-ml-sm">Bu kullanıcıyı silmek istediğinize emin misiniz?</span>
        </q-card-section>

        <q-card-actions align="right">
          <q-btn flat label="İptal" color="primary" v-close-popup />
          <q-btn flat label="Sil" color="negative" @click="deleteUser" />
        </q-card-actions>
      </q-card>
    </q-dialog>
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
    const dialogVisible = ref(false)
    const selectedUserId = ref(null)

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
        console.error('API Hatası:', error)
      }
    }

    onMounted(fetchClients)

    // Kullanıcı silme fonksiyonu
    const deleteUser = async () => {
      try {
        await axios.delete(`http://localhost:9000/delete/${selectedUserId.value}`, {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`,
          },
        })
        $q.notify({
          type: 'positive',
          message: 'Kullanıcı başarıyla silindi!',
        })
        // Kullanıcı listesini güncelle
        fetchClients()
      } catch (error) {
        $q.notify({
          type: 'negative',
          message: `Kullanıcı silinirken hata oluştu: ${error.message}`,
        })
        console.error('Silme Hatası:', error)
      } finally {
        dialogVisible.value = false
        selectedUserId.value = null
      }
    }

    // Silme işlemi için onay diyaloğunu açma fonksiyonu
    const confirmDelete = (id) => {
      selectedUserId.value = id
      dialogVisible.value = true
    }

    // Tablo Sütunları
    const columns = [
      { name: 'id', label: 'ID', align: 'left', field: (row) => row.id, sortable: true },
      { name: 'name', label: 'Adı', align: 'left', field: 'name', sortable: true },
      { name: 'email', label: 'E-Posta', align: 'left', field: 'email' },
      { name: 'role', label: 'Rol', align: 'center', field: 'role', sortable: true },
      { name: 'actions', label: 'İşlemler', align: 'center' }
    ]

    return { clients, columns, confirmDelete, deleteUser, dialogVisible }
  },
}
</script>
