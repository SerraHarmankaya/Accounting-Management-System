<template>
  <q-layout view="lHh Lpr lFf">
    <q-page-container>
      <q-page class="flex flex-center bg-image">
        <q-card class="q-pa-md" style="width: 400px">
          <q-card-section>
            <div class="text-h6">Kayıt Ol</div>
          </q-card-section>

          <q-card-section>
            <p v-if="message">{{ message }}</p>
            <q-input v-model="name" placeholder="İsim Soyisim" outlined />
            <q-input v-model="email" placeholder="E-posta" outlined class="q-mt-md" />
            <q-input
              v-model="password"
              placeholder="Şifre"
              type="password"
              outlined
              class="q-mt-md"
            />
            <q-select v-model="role" :options="options" label="Role" />
          </q-card-section>

          <q-card-actions align="right">
            <q-btn label="Kayıt Ol" color="primary" @click="createUser" />
          </q-card-actions>
        </q-card>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script>
import { ref } from 'vue'
//import { useRouter } from 'vue-router'
import axios from 'axios'

export default {
  data() {
    // const router = useRouter()

    // const goToLogin = () => {
    //   router.push('/login') // Login sayfasına yönlendir
    // }

    return {
      message: '',
      name: '',
      email: '',
      password: '',
      role: '', // Seçilen değeri tutar
      options: ref(['User', 'Admin']),
      //goToLogin,
    }
  },
  methods: {
    async createUser() {
      try {
        const userData = {
          name: this.name.trim(),
          email: this.email.trim(), // Boşlukları temizliyoruz
          password: this.password.trim(),
          role: this.role.toLowerCase(),
        }

        console.log('Gönderilen JSON:', JSON.stringify(userData)) // Konsolda JSON verisini göster

        const response = await axios.post('http://localhost:9000/register', userData, {
          headers: { 'Content-Type': 'application/json' },
        })

        if (response.data.success) {
          this.$q.notify({ type: 'positive', message: 'User saved successfully!' })
          console.log('succeded!')

          this.$router.push('/login')
        } else {
          this.$q.notify({ type: 'negative', message: 'Failed to save user!' })
          console.log('failed!')
        }
      } catch (error) {
        console.error('Create user error:', error.response ? error.response.data : error.message)
      }
    },
  },
}
</script>
<style>
.bg-image {
  background-image: linear-gradient(135deg, #650606ee 0%, #dba6be 100%);
}
</style>
