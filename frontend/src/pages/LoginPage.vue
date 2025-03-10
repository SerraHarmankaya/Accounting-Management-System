<template>
  <q-layout view="lHh Lpr lFf">
    <q-page-container>
      <q-page class="flex flex-center bg-image">
        <q-card class="q-pa-md" style="width: 400px">
          <q-card-section>
            <div class="text-h6">Giriş Yap</div>
          </q-card-section>

          <q-card-section>
            <q-input filled v-model="email" label="E-posta" type="email" />
            <q-input filled v-model="password" label="Şifre" type="password" class="q-mt-md" />

            <div>
              Hesabınız yok mu?
              <q-btn flat dense label="Kaydolun" color="primary" @click="goToRegister" />
            </div>
          </q-card-section>

          <q-card-actions class="justify-end">
            <q-btn label="Giriş Yap" color="primary" @click="loginUser" />
          </q-card-actions>
        </q-card>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script>
import axios from 'axios'

//import { errorMessages } from 'vue/compiler-sfc'

export default {
  data() {
    return {
      message: '',
      email: '',
      password: '',
    }
  },
  methods: {
    async loginUser() {
      try {
        const userData = {
          email: this.email.trim(),
          password: this.password.trim(),
        }

        const response = await axios.post('http://localhost:9000/login', userData, {
          headers: { 'Content-Type': 'application/json' },
        })

        if (response.data.success) {
          localStorage.setItem('token', response.data.token)
          window.dispatchEvent(new Event('storage'))

          this.$q.notify({ type: 'positive', message: 'User Logged in successfully!' })
          console.log('succeded!')

          if (response.data.role == 'admin') {
            this.$router.push('/admin-page')
          } else {
            this.$router.push('/user-page')
          }
        } else {
          this.$q.notify({ type: 'negative', message: 'Failed to Login!' })
          console.log('failed!')
        }
      } catch (error) {
        console.log('Create user error:', error.response ? error.response.data : error.message)
      }
    },

    goToRegister() {
      this.$router.push('/register')
    },
  },
}
</script>
<style>
.bg-image {
  background-image: linear-gradient(135deg, #650606ee 0%, #dba6be 100%);
}
</style>
