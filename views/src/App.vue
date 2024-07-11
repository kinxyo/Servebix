<script setup>
import DashboardControl from "@/components/fixed/dashboard_controller.vue";
import PatientCard from "@/components/fixed/patientCard.vue";
import {onBeforeMount, ref} from "vue";
import {ShowAllPatients} from "@/composables/dal.js";

const records = ref([]);

onBeforeMount(async () => {
  records.value = await ShowAllPatients()
})

</script>

<template>
  <main class="dark flex flex-col gap-6">

    <DashboardControl/>

    <div class="flex flex-col xl:grid xl:grid-cols-6 gap-8 px-12">
      <PatientCard
          class="transition hover:scale-110 hover:shadow-sm hover:shadow-emerald-700 cursor-pointer"
          v-for="card in records"
          :key="card.id"
          :id="card.id"
          :name="card.name"
          :email="card.email"
          :phone="card.phone"
          :password="card.password"
          :createdAt="new Date(card.createdAt).toLocaleDateString()"
      />
    </div>

  </main>
</template>