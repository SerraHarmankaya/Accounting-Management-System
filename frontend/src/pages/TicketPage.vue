<template>
  <q-page class="flex flex-center">
    <q-card class="q-pa-md text-center" style="width: 400px">
      <q-card-section>
        <div class="text-h5">Ticket Oluşturun</div>
        <q-input v-model="title" label="Title" outlined class="q-mt-md" />
        <q-input v-model="content" label="Content" outlined class="q-mt-md" />
      </q-card-section>

      <q-card-actions align="right">
        <q-btn label="Kaydet" color="primary" @click="createTicket" />
      </q-card-actions>
    </q-card>
  </q-page>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      title: '', // Başlık alanı
      content: '', // İçerik alanı
    }
  },
  methods: {
    async createTicket() {
      try {
        const token = localStorage.getItem('token') // Kullanıcının giriş yaptığı token

        if (!token) {
          this.$q.notify({ type: 'negative', message: 'Giriş yapmalısınız!' })
          this.$router.push('/login') // Eğer token yoksa login sayfasına yönlendir
          return
        }

        const ticketData = {
          title: this.title.trim(),
          content: this.content.trim(),
        }
        console.log(ticketData)
        const response = await axios.post('http://localhost:9000/create/user-tickets', ticketData, {
          headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${token}`, // JWT Token ekleniyor
          },
        })

        if (response.data.success) {
          this.$q.notify({ type: 'positive', message: 'Ticket başarıyla oluşturuldu!' })
          this.$router.push('/user/tickets') // Ticket oluşturulduktan sonra yönlendirme
        } else {
          this.$q.notify({ type: 'negative', message: 'Ticket oluşturulamadı!' })
        }
      } catch (error) {
        console.error(
          'Ticket oluşturma hatası:',
          error.response ? error.response.data : error.message,
        )
        this.$q.notify({ type: 'negative', message: 'Bir hata oluştu!' })
      }
    },
  },
}
</script>
