<script setup>
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card/index.js";
import {Button} from "@/components/ui/button/index.js";


defineProps(['title', 'subtitle', 'content', 'action'])

</script>


<template>

  <Card>
    <CardHeader>
      <CardTitle>{{ title }}</CardTitle>
      <CardDescription>{{ subtitle }}</CardDescription>
    </CardHeader>
    <CardContent> "{{ content }}"</CardContent>
    <CardFooter>
      <Button @click="$emit('greet')" > {{ action }} </Button>
    </CardFooter>
  </Card>

</template>

<style scoped>

</style>