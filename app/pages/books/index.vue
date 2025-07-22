<script setup lang="ts">
const { getBooks } = useBooks()

const { data } = await getBooks()
console.log(data)
</script>

<template>
  <div class="flex flex-col gap-4 items-center text-red-500 text-2xl">
    <h1>Books</h1>
    <ul>
      <li v-for="book in data" :key="book.id">
        {{ book.title }}
      </li>
    </ul>
  </div>
</template>
  