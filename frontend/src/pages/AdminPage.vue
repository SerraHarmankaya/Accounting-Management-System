<template>
  <q-page class="flex flex-center">
    <div class="column q-gutter-md">
      <q-btn @click="listAllTickets" push color="grey-6" label="LIST ALL TICKETS" />
      <q-btn @click="exportToExcel" push color="green" label="EXPORT TO EXCEL" icon="save_alt" />
      <q-table
        class="q-mt-md fixed-table"
        overflow: auto
        flat
        bordered
        title="Tickets"
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
      </q-table>
    </div>
  </q-page>
</template>

<script>
import axios from 'axios'
import * as XLSX from 'xlsx' // SheetJS kütüphanesini içe aktar

export default {
  data() {
    return {
      columns: [
        {
          name: 'id',
          label: 'ID',
          align: 'left',
          field: (row) => row.id,
          sortable: true,
        },
        { name: 'title', label: 'Title', align: 'left', field: 'title', sortable: true },
        { name: 'content', label: 'Content', align: 'left', field: 'content' },
        {
          name: 'status',
          label: 'Status',
          align: 'center',
          field: 'status',
          sortable: true,
        },
        { name: 'created_by', label: 'Created By', align: 'left', field: 'created_by' },
      ],
      message: '',
      title: '',
      content: '',
      status: '',
      tickets: [],
    }
  },
  methods: {
    async listAllTickets() {
      try {
        const token = localStorage.getItem('token') // Kullanıcının giriş yaparken aldığı token
        console.log('token: ', token)
        if (!token) {
          console.log('Kullanıcı oturum açmamış, login sayfasına yönlendiriliyor...')
          this.$router.push('/login') // Kullanıcıyı giriş sayfasına yönlendir
          return
        }
        const response = await axios.get('http://localhost:9000/admin-page', {
          headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${token}`, // JWT token ekleniyor
          },
        })

        console.log('API Response:', response.data)
        this.tickets = response.data
      } catch (error) {
        console.error(
          'Ticketları listeleme hatası:',
          error.response ? error.response.data : error.message,
          console.error('Hata Detayı:', error.response ? error.response.data : error.message),
        )
      }
    },

    // Excel çıktısı alma fonksiyonu
    exportToExcel() {
      if (this.tickets.length === 0) {
        this.$q.notify({ type: 'negative', message: 'Hiç ticket verisi bulunamadı.' })
        console.warn('Hiç ticket verisi bulunamadı.')
        return
      }

      // JSON verisini Excel formatına dönüştür
      const worksheet = XLSX.utils.json_to_sheet(this.tickets)

      // Yeni bir çalışma kitabı (Workbook) oluştur
      const workbook = XLSX.utils.book_new()
      XLSX.utils.book_append_sheet(workbook, worksheet, 'Tickets')

      // Excel dosyasını indir
      XLSX.writeFile(workbook, 'ticket-listesi.xlsx')

      this.$q.notify({ type: 'positive', message: 'Excel dosyası başarıyla indirildi!' })
      console.log('Excel dosyası başarıyla indirildi!')
    },

    getStatusColor(status) {
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
    },
  },
}
</script>
<style>
.fixed-table {
  max-width: 800px;
  min-width: 600px;
  height: 400px;
  overflow: auto; /* Taşan içerikler kaydırılabilir olsun */
}
</style>
