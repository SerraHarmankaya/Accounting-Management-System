<template>
  <q-layout view="lHh Lpr lFf">
    <q-drawer v-model="leftDrawerOpen" show-if-above bordered>
      <q-list>
        <q-item-label class="text-h6" header>
          <q-img src="mini.png" width="120px" class="q-mb-lg" />
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

<script>
export default {
  data() {
    return {
      leftDrawerOpen: true,
      linksList: [
        {
          title: 'Islemler',
          icon: 'work',
          link: 'https://quasar.dev',
          children: [
            { title: 'Sub 1', link: '/operations/sub1', icon: 'list' },
            { title: 'Sub 2', link: '/operations/sub2', icon: 'list' },
          ],
        },
        {
          title: 'Github',
          icon: 'code',
          link: 'https://github.com/quasarframework',
        },
        {
          title: 'Community',
          icon: 'people',
          children: [
            { title: 'Discord', link: 'https://chat.quasar.dev', icon: 'chat' },
            {
              title: 'Forum',
              link: 'https://forum.quasar.dev',
              icon: 'record_voice_over',
            },
          ],
        },
        {
          title: 'Social Media',
          icon: 'rss_feed',
          children: [
            { title: 'Twitter', link: 'https://twitter.quasar.dev', icon: 'rss_feed' },
            { title: 'Facebook', link: 'https://facebook.quasar.dev', icon: 'public' },
          ],
        },
        {
          title: 'Settings',
          icon: 'settings',
          link: '/settings',
        },
      ],
    }
  },
}
</script>
