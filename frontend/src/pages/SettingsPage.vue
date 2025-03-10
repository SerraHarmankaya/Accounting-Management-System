<template>
  <q-page class="q-pa-md">
    <q-card class="q-pa-md" style="max-width: 500px; margin: auto">
      <q-card-section>
        <div class="text-h5 text-primary">Ayarlar</div>
      </q-card-section>

      <q-separator />

      <!-- Kullanıcı Bilgileri -->
      <q-card-section>
        <q-input filled v-model="user.name" label="Ad Soyad" />
        <q-input filled v-model="user.email" label="E-posta" type="email" class="q-mt-md" />
      </q-card-section>

      <q-separator />

      <!-- Tema Seçimi -->
      <q-card-section>
        <q-item>
          <q-item-section>
            <q-item-label>Tema</q-item-label>
          </q-item-section>
          <q-item-section side>
            <q-toggle v-model="darkMode" label="Karanlık Mod" @update:model-value="toggleTheme" />
          </q-item-section>
        </q-item>
      </q-card-section>

      <q-separator />

      <!-- Bildirim Seçenekleri -->
      <q-card-section>
        <q-item>
          <q-item-section>
            <q-item-label>Bildirimler</q-item-label>
          </q-item-section>
          <q-item-section side>
            <q-toggle v-model="user.notifications" />
          </q-item-section>
        </q-item>
      </q-card-section>

      <q-separator />

      <!-- Şifre Değiştir -->
      <q-card-section class="text-center">
        <q-btn color="primary" label="Şifreyi Değiştir" @click="changePassword" />
      </q-card-section>

      <q-separator />

      <!-- Kaydet ve Vazgeç Butonları -->
      <q-card-actions align="right">
        <q-btn flat label="Vazgeç" color="negative" @click="resetSettings" />
        <q-btn label="Kaydet" color="primary" @click="saveSettings" />
      </q-card-actions>
    </q-card>
  </q-page>
</template>

<script setup>
import { ref } from 'vue'
import { useQuasar } from 'quasar'

// Quasar objesi
const $q = useQuasar()

// Kullanıcı bilgileri (örnek veriler)
const user = ref({
  name: 'Ahmet Yılmaz',
  email: 'ahmet@example.com',
  notifications: true,
})

// Karanlık mod kontrolü
const darkMode = ref($q.dark.isActive)

// Tema değiştirme fonksiyonu
const toggleTheme = () => {
  $q.dark.set(darkMode.value)
}

// Şifre değiştirme
const changePassword = () => {
  $q.notify({
    type: 'info',
    message: 'Şifre değiştirme sayfasına yönlendiriliyorsunuz...',
  })
}

// Ayarları kaydetme
const saveSettings = () => {
  $q.notify({
    type: 'positive',
    message: 'Ayarlar başarıyla kaydedildi!',
  })
}

// Değişiklikleri iptal etme (resetleme)
const resetSettings = () => {
  user.value = {
    name: 'Ahmet Yılmaz',
    email: 'ahmet@example.com',
    notifications: true,
  }
  darkMode.value = $q.dark.isActive
  $q.notify({
    type: 'negative',
    message: 'Değişiklikler iptal edildi.',
  })
}
</script>
