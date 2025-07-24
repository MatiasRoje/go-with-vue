<script setup lang="ts">
const route = useRoute()
const slugParam = route.params.slug
const slug = Array.isArray(slugParam) ? slugParam[0] : slugParam

if (!slug) {
  navigateTo('/books')
  throw new Error("No slug provided")
}

const { getBook } = useBooks()

const { data } = await getBook(slug as string)

const toast = useToast()

function showToast() {
  toast.add({
    title: 'Sleep',
    description: 'We need to sleep',
    icon: 'i-lucide-bed',
    ui: {
      description: "text-2xl text-red-500 font-bold"  
    }
  })
}
</script>

<template>
  <main class="flex flex-col gap-4 items-center my-4">
    <h1 v-if="data">{{ data?.title }}</h1>
    <UButton @click="showToast">Click me</UButton>
  </main>
</template>