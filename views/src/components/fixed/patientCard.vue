<script setup>
import {Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle} from "@/components/ui/card/index.js";

const props = defineProps({
  id: "",
  name: "",
  email: "",
  phone: "",
  password: "",
  createdAt: ""
})


</script>

<template>

  <Card>
    <CardHeader>
      <CardTitle> {{ name }}</CardTitle>
      <CardDescription>  {{ id }} </CardDescription>
    </CardHeader>
    <CardContent class="text-emerald-500 font-bold"> {{ email }} </CardContent>
    <CardFooter class="text-slate-400">
<!--      <p> {{ phone.String }} </p>-->
      <p> {{ createdAt }} </p>
    </CardFooter>
  </Card>

</template>

<style scoped>

</style>