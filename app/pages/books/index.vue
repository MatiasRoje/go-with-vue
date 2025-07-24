<script setup lang="ts">
const { getBooks } = useBooks()

const { data } = await getBooks()
</script>

<template>
  <main class="flex flex-col gap-4 items-center my-4">
    <ul class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <li v-for="book in data" :key="book.id">
        <ULink :href="`/books/${book.slug}`">
          <UCard
            class="border border-gray-200 overflow-hidden flex flex-col h-full hover:cursor-pointer"
            :ui="{
              header: 'p-0 sm:px-0',
              body: 'p-3'
            }"
          >
            <template #header>
              <nuxt-img
                :src="`http://localhost:3001/static/covers/${book.slug}.jpg`"
                class="w-full min-h-64 object-cover"
              />
            </template>
            <div class="flex flex-col flex-1">
              <h3 class="text-lg font-bold mb-1">{{ book.title }}</h3>
              <p class="text-sm text-gray-600 mb-1">By {{ book.author.author_name }}</p>
              <p class="text-xs text-gray-400 mb-2">{{ book.publication_year }}</p>
              <div class="flex flex-wrap gap-2 mt-auto">
                <span
                  v-for="genre in book.genres"
                  :key="genre.id"
                  class="bg-indigo-100 text-indigo-800 rounded-full px-2 py-0.5 text-xs text-center font-semibold mt-2"
                >
                  {{ genre.genre_name }}
                </span>
              </div>
            </div>
          </UCard>        
        </ULink>
      </li>
    </ul>
  </main>
</template>