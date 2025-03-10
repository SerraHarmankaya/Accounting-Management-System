<template>
  <q-page class="flex flex-center">
    <q-header> </q-header>
    <q-card class="my-card">
      <q-card-section class="bg-primary text-white">
        <div class="text-h6">Yapmak istediginiz islemi secin</div>
      </q-card-section>
      <q-card-actions vertical align="center">
        <q-btn flat @click="goToCreateTicket">Ticket Olustur</q-btn>
        <q-btn flat @click="goToViewTickets">Ticketlarimi goruntule</q-btn>
      </q-card-actions></q-card
    >
  </q-page>
</template>

<script>
export default {
  methods: {
    goToCreateTicket() {
      this.$router.push('/create/user-tickets') // Sayfa yönlendirme
    },
    goToViewTickets() {
      this.$router.push('/user/tickets') // Sayfa yönlendirme
    },
  },
}
</script>
