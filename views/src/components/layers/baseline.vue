<script setup>
import GreetMe from "@/components/fixed/dashboard_controller.vue";
import PatientCard from "@/components/fixed/patientCard.vue";
import {onBeforeMount, ref} from "vue";
import {ShowAllPatients} from "@/composables/dal.js";

const records = ref([])

async function setPatientCard() {
  const data = await ShowAllPatients();

  console.log(data);

  records.value = await data;

}

onBeforeMount(async () => {
  await setPatientCard();
})

</script>

<template>

  <main class="flex flex-col gap-6">



  </main>

</template>

<style scoped>

</style>