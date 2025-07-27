<script setup lang="ts">
const router = useRouter()
const route = useRoute()
const slugParam = route.params.slug
const slug = Array.isArray(slugParam) ? slugParam[0] : slugParam

if (!slug) {
  navigateTo('/books')
  throw new Error("No slug provided")
}

const { getBook } = useBooks();

const { loading, error, data } = await getBook(slug as string);

</script>

<template>
  <main class="max-w-3xl mx-auto py-12 px-4">
    <!-- Loading State -->
    <div v-if="loading" class="flex gap-8 mb-10">
      <div>
        <USkeleton class="w-64 h-96 rounded mb-2" />
      </div>
      <div class="flex-1 flex flex-col justify-between">
        <USkeleton class="h-8 w-3/4 mb-4 rounded" />
        <USkeleton class="h-5 w-1/2 mb-6 rounded" />
        <USkeleton class="h-48 w-80 mb-2 rounded" />
        <div class="flex gap-2 mt-auto">
          <USkeleton class="h-6 w-14 rounded-full" />
          <USkeleton class="h-6 w-14 rounded-full" />
        </div>
      </div>
    </div>

    <!-- Error State -->
    <p v-else-if="error" class="mb-10">Something went wrong. Please try again later.</p>

    <!-- Data State -->
    <div v-else-if="data" class="flex gap-8 mb-10">
      <div>
        <nuxt-img :src="`http://localhost:3001/static/covers/${data.slug}.jpg`" class="w-64 h-full  max-w-64" />
      </div>
      <div class="flex flex-col justify-between">
        <h1 class="text-3xl font-bold mb-1">{{ data.title }}</h1>
        <div class="mb-6 text-gray-700">
          By {{ data.author.author_name }}
        </div>
        <div class="min-w-72">
          <p class="text-base text-gray-800 text-justify">{{ data.description }}</p>
        </div>
        <div class="flex flex-wrap gap-2 mt-auto">
          <span
            v-for="genre in data.genres"
            :key="genre.id"
            class="bg-indigo-100 text-indigo-800 rounded-full px-2 py-1 text-xs text-center font-semibold mt-2"
          >
            {{ genre.genre_name }}
          </span>
        </div>
      </div>
    </div>

    <UButton class="hover:cursor-pointer" @click="router.back()">
      ← Back to Books
    </UButton>
  </main>
</template>