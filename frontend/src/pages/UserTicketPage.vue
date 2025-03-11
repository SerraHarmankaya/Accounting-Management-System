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

    </q-table>
  </q-page>
</template>

<script>
import { ref, onMounted } from "vue";
import axios from "axios";
import * as XLSX from "xlsx"; // SheetJS kütüphanesi
import { useQuasar } from "quasar";

export default {
  setup() {
    const $q = useQuasar();
    const tickets = ref([]);
    const dialogVisible = ref(false);
    const selectedTicketId = ref(null);

    // Kullanıcının Ticket'larını Getir
    const fetchUserTickets = async () => {
      try {
        const token = localStorage.getItem("token");
        console.log("Token:", token);

        if (!token) {
          console.log("Kullanıcı oturum açmamış, login sayfasına yönlendiriliyor...");
          window.location.href = "/login";
          return;
        }

        const response = await axios.get("http://localhost:9000/user/tickets", {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        if (response.status === 200) {
          tickets.value = response.data;
          console.log("Kullanıcının ticket'ları:", tickets.value);
        } else {
          $q.notify({ type: "negative", message: "Ticketlar alınamadı!" });
        }
      } catch (error) {
        console.error(
          "Ticketları alma hatası:",
          error.response ? error.response.data : error.message
        );
        $q.notify({ type: "negative", message: "Bir hata oluştu!" });
      }
    };

    // Excel Çıktısı Alma Fonksiyonu
    const exportToExcel = () => {
      if (tickets.value.length === 0) {
        $q.notify({ type: "negative", message: "Hiç ticket verisi bulunamadı." });
        console.warn("Hiç ticket verisi bulunamadı.");
        return;
      }

      // JSON verisini Excel formatına dönüştür
      const worksheet = XLSX.utils.json_to_sheet(tickets.value);

      // Yeni bir çalışma kitabı oluştur
      const workbook = XLSX.utils.book_new();
      XLSX.utils.book_append_sheet(workbook, worksheet, "Tickets");

      // Excel dosyasını indir
      XLSX.writeFile(workbook, "ticket-listesi.xlsx");

      $q.notify({ type: "positive", message: "Excel dosyası başarıyla indirildi!" });
      console.log("Excel dosyası başarıyla indirildi!");
    };

    // Ticket Silme Fonksiyonu
    const deleteByID = async () => {
      try {
        await axios.delete(`http://localhost:9000/delete/${selectedTicketId.value}`, {
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
          },
        });

        $q.notify({ type: "positive", message: "Kullanıcı başarıyla silindi!" });

        // Listeyi güncelle
        fetchUserTickets();
      } catch (error) {
        $q.notify({
          type: "negative",
          message: `Kullanıcı silinirken hata oluştu: ${error.message}`,
        });
        console.error("Silme Hatası:", error);
      } finally {
        dialogVisible.value = false;
        selectedTicketId.value = null;
      }
    };

    // Status'e göre renk belirleyen fonksiyon
    const getStatusColor = (status) => {
      if (!status) return "grey"; // Null veya undefined gelirse
      switch (status.toLowerCase()) {
        case "open":
          return "blue";
        case "ongoing":
          return "orange";
        case "closed":
          return "green";
        case "pending":
          return "yellow";
        case "canceled":
          return "red";
        default:
          return "grey";
      }
    };

    // Silme işlemi için onay diyaloğunu açma fonksiyonu
    const confirmDelete = (id) => {
      selectedTicketId.value = id;
      dialogVisible.value = true;
    };

    // Tablo Sütunları
    const columns = [
      { name: "id", label: "ID", align: "left", field: (row) => row.id, sortable: true },
      { name: "title", label: "TITLE", align: "left", field: "title", sortable: true },
      { name: "content", label: "CONTENT", align: "left", field: "content" },
      { name: "status", label: "STATUS", align: "center", field: "status", sortable: true },
      { name: "actions", label: "DELETE", align: "center" },
    ];

    // Sayfa yüklendiğinde verileri çek
    onMounted(fetchUserTickets);

    return { tickets, columns, fetchUserTickets, exportToExcel, confirmDelete, deleteByID, dialogVisible, getStatusColor };
  },
};
</script>

