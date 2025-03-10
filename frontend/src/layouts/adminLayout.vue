<template>
  <q-layout view="lHh Lpr lFf">
    <q-header v-if="isAuthenticated" elevated>
      <q-toolbar>
        <q-btn flat dense round icon="menu" aria-label="Menu" @click="toggleLeftDrawer" />
        <q-toolbar-title> Ticket </q-toolbar-title>
        <q-space />
        <q-btn flat no-caps icon="person">
          <!--  <span class="q-ml-sm">{{ userName }}</span> -->

          <span class="q-ml-sm">Serra Harmankaya</span>

          <q-menu fit anchor="bottom right" self="top right">
            <q-list style="min-width: 150px">
              <q-item clickable v-close-popup @click="goToProfile">
                <q-item-section>👤 Profile</q-item-section>
              </q-item>
              <q-item clickable v-close-popup @click="logout">
                <q-item-section>🚪 Logout</q-item-section>
              </q-item>
            </q-list>
          </q-menu>
        </q-btn>
      </q-toolbar>
    </q-header>

    <q-drawer v-model="leftDrawerOpen" show-if-above bordered>
      <q-list>
        <q-item-label class="text-h6" header>
          <q-img src="apple.png" width="120px" class="q-mb-lg" />
          <br />
          Essential Links
          <br />
        </q-item-label>

        <!-- Menü elemanlarını oluştur -->
        <template v-for="link in linksList" :key="link.title">
          <!-- Eğer alt link yoksa normal link olarak göster -->
          <q-item v-if="!link.children" :to="link.link" clickable>
            <q-item-section avatar v-if="link.icon">
              <q-icon :name="link.icon" />
            </q-item-section>
            <q-item-section>
              <q-item-label>{{ link.title }}</q-item-label>
            </q-item-section>
          </q-item>

          <!-- Eğer alt başlıkları varsa, açılabilir bir menü yap -->
          <q-expansion-item v-else :icon="link.icon" :label="link.title">
            <q-list>
              <q-item v-for="sub in link.children" :key="sub.title" :to="sub.link" clickable>
                <q-item-section avatar v-if="sub.icon">
                  <q-icon :name="sub.icon" />
                </q-item-section>
                <q-item-section>
                  <q-item-label>{{ sub.title }}</q-item-label>
                </q-item-section>
              </q-item>
            </q-list>
          </q-expansion-item>
        </template>
      </q-list>
    </q-drawer>
    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
//import EssentialLink from 'components/EssentialLink.vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

// Reactive state for authentication and drawer
const isAuthenticated = ref(false)
const leftDrawerOpen = ref(false)
const router = useRouter()

// Check authentication status
async function checkAuthentication() {
  try {
    const token = localStorage.getItem('token')
    if (!token) {
      console.log('Token bulunamadı, kullanıcı giriş yapmamış.')
      isAuthenticated.value = false
      return
    }

    const response = await axios.get('http://localhost:9000/auth-check', {
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
    })

    if (response.data.success) {
      console.log('Kullanıcı doğrulandı:', response.data)
      isAuthenticated.value = true
    } else {
      console.log('Doğrulama başarısız.')
      isAuthenticated.value = false
    }
  } catch (error) {
    console.log('Kimlik doğrulama hatası:', error.response ? error.response.data : error.message)
    isAuthenticated.value = false
  }
}

// Watch for changes to the token in localStorage
watch(
  () => localStorage.getItem('token'),
  async (newToken, oldToken) => {
    if (newToken !== oldToken) {
      await checkAuthentication()
    }
  },
)

// Check authentication on component mount
onMounted(() => {
  checkAuthentication()
})

function logout() {
  localStorage.removeItem('token')
  isAuthenticated.value = false // Manually update the state
  router.push('/login') // Yönlendirme yapılacak sayfa
}

// Toggle drawer
function toggleLeftDrawer() {
  leftDrawerOpen.value = !leftDrawerOpen.value
}

// const options = [
//   {
//     title: 'Profile',
//     icon: 'person',
//     //link: '/profile',
//   },
//   {
//     title: 'Logout',
//     icon: 'logout',
//     //link: '/logout',
//   },
// ]

// Links for the drawer
const linksList = [
  {
    title: 'Islemler',
    icon: 'work',
    link: 'https://quasar.dev',
  },
  {
    title: 'Client',
    icon: 'groups',
    link: '/clients',
    children: [
      { title: 'List Clients', link: '/clients/list', icon: 'groups' },
      { title: 'Sub 2', link: '/settings/sub2', icon: 'groups' },
    ],
  },
  {
    title: 'Discord Chat Channel',
    icon: 'chat',
    link: 'https://chat.quasar.dev',
  },
  {
    title: 'Forum',
    icon: 'record_voice_over',
    link: 'https://forum.quasar.dev',
  },
  {
    title: 'Twitter',
    icon: 'rss_feed',
    link: 'https://twitter.quasar.dev',
  },
  {
    title: 'Facebook',
    icon: 'public',
    link: 'https://facebook.quasar.dev',
  },
  {
    title: 'Settings',
    icon: 'settings',
    link: '/admin/settings',
  },
]
</script>
