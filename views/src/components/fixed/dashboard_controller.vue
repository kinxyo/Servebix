<script setup>

import {Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle} from "@/components/ui/card/index.js";
import {Button} from "@/components/ui/button/index.js";
import {onBeforeMount, ref} from "vue";
import {fetchGreet} from "@/composables/dal.js";
import {refreshPage} from "@/utils/utils.js";

const greet = ref("");

onBeforeMount(async () => {
  greet.value = await fetchGreet()
});

</script>

<template>

  <Card>
    <CardHeader>
      <CardTitle> Server Dashboard</CardTitle>
      <CardDescription> © 2024 Ebix Inc.</CardDescription>
    </CardHeader>
    <CardContent class="text-emerald-500 font-bold">{{ greet }}</CardContent>
    <CardFooter>
      <Button @click="refreshPage"> Update Status </Button>
    </CardFooter>
  </Card>

</template>

<style scoped>

</style>