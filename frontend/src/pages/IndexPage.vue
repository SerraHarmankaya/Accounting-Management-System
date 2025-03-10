<template>
  <q-layout view="lHh Lpr lFf">
    <q-page-container>
      <q-page class="flex flex-center bg-image">
        <q-card class="q-pa-md text-center" style="width: 400px">
          <q-card-section>
            <div class="text-h5">Hoş Geldiniz</div>
            <div class="text-caption q-mt-sm">Lütfen giriş yapın veya kaydolun.</div>
          </q-card-section>

          <q-card-actions vertical class="q-mt-md">
            <q-btn label="Giriş Yap" color="primary" @click="goToLogin" class="full-width" />
            <q-btn
              label="Kayıt Ol"
              color="accent"
              @click="goToRegister"
              class="full-width q-mt-sm"
            />
          </q-card-actions>
        </q-card>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script>
import { useRouter } from 'vue-router'

export default {
  setup() {
    const router = useRouter()

    const goToLogin = () => {
      router.push('/login') // Login sayfasına yönlendir
    }

    const goToRegister = () => {
      router.push('/register') // Register sayfasına yönlendir
    }

    return { goToLogin, goToRegister }
  },
}
</script>

<style scoped>
.full-width {
  width: 100%;
}
.bg-image {
  background-image: linear-gradient(135deg, #650606ee 0%, #dba6be 100%);
}
</style>
