<script setup>

import {Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle} from "@/components/ui/card/index.js";
import {Button} from "@/components/ui/button/index.js";
import {onBeforeMount, ref} from "vue";
import {fetchGreet} from "@/composables/dal.js";

const greet = ref("");

onBeforeMount(async () => {
  await setGreet();
});

async function setGreet() {
  greet.value = await fetchGreet()
}

</script>

<template>

  <Card>
  <CardHeader>
    <CardTitle> Server Dashboard </CardTitle>
    <CardDescription> © 2024 Ebix Inc. </CardDescription>
  </CardHeader>
  <CardContent class="text-emerald-500 font-bold">{{ greet }}</CardContent>
  <CardFooter>
    <Button @click="setGreet" > Fetch Status </Button>
  </CardFooter>
  </Card>

</template>

<style scoped>

</style>