<template>
  <q-page padding>
    <q-table
      flat
      bordered
      title="Ticket Listesi"
      :rows="tickets"
      :columns="columns"
      row-key="id"

    >
    <template v-slot:body-cell-status="props">
          <q-td :props="props">
            <q-chip :color="getStatusColor(props.row.status)" text-color="white">
              {{ props.row.status }}
            </q-chip>
          </q-td>
        </template>

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
          <q-avatar icon="warning" color="red" text-color="white" />
          <span class="q-ml-sm">Bu ticket'i silmek istediğinize emin misiniz?</span>
        </q-card-section>

        <q-card-actions align="right">
          <q-btn flat label="İptal" color="primary" v-close-popup />
          <q-btn flat label="Sil" color="negative" @click="deleteByID" />
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
    const tickets = ref([])
    const dialogVisible = ref(false)
    const selectedTicketId = ref(null)

    // API'den veri çekme fonksiyonu
    const fetchTickets = async () => {
      try {
        const response = await axios.get('http://localhost:9000/clients/ticket', {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`,
          },
        })

        
        tickets.value = response.data
      } catch (error) {
        $q.notify({
          type: 'negative',
          message: 'Veriler alınırken hata oluştu!',
          
        })
        console.error('API Hatası:', error)
      }
    }
    

    onMounted(fetchTickets)

    // Kullanıcı silme fonksiyonu
    const deleteByID = async () => {
      try {
        await axios.delete(`http://localhost:9000/delete/${selectedTicketId.value}`, {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`,
          },
        })
        $q.notify({
          type: 'positive',
          message: 'Kullanıcı başarıyla silindi!',
        })
        // Kullanıcı listesini güncelle
        fetchTickets()
      } catch (error) {
        $q.notify({
          type: 'negative',
          message: `Kullanıcı silinirken hata oluştu: ${error.message}`,
        })
        console.error('Silme Hatası:', error)
      } finally {
        dialogVisible.value = false
        selectedTicketId.value = null
      }
    }

      // Status'e göre renk belirleyen fonksiyon
      const getStatusColor = (status) => {
      if (!status) return 'grey' // Null veya undefined gelirse
      switch (status.toLowerCase()) {
        case 'open':
          return 'blue'
        case 'in progress':
          return 'orange'
        case 'closed':
          return 'green'
        default:
          return 'grey'
      }
    }

    // Silme işlemi için onay diyaloğunu açma fonksiyonu
    const confirmDelete = (id) => {
      selectedTicketId.value = id
      dialogVisible.value = true
    }

    // Tablo Sütunları
    const columns = [
      { name: 'id', label: 'ID', align: 'left', field: (row) => row.id, sortable: true },
      { name: 'title', label: 'TITLE', align: 'left', field: 'title', sortable: true },
      { name: 'content', label: 'CONTENT', align: 'left', field: 'content' },
      { name: 'status', label: 'STATUS', align: 'center', field: 'status', sortable: true },
      { name: 'created_by', label: 'CREATED BY', align: 'center', field: 'created_by', sortable: true },
      { name: 'actions', label: 'DELETE', align: 'center' }
    ]

    return { tickets, columns, confirmDelete, deleteByID, dialogVisible, getStatusColor }
  },
}
</script>
