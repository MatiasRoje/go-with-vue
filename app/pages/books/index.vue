<script setup lang="ts">
const router = useRouter()
const route = useRoute()
const { getBooks } = useBooks()
const { error, data, loading } = await getBooks()

// Get unique genres from data, flattening all genres arrays
const genres = computed(() => {
  if (!data) return []
  const allGenres = data.flatMap(book => book.genres || [])
  // Remove duplicates by id
  const unique = []
  const seen = new Set()
  for (const genre of allGenres) {
    if (!seen.has(genre.id)) {
      seen.add(genre.id)
      unique.push(genre)
    }
  }
  
  return unique
})

const selectedGenre = ref<{label: string, value: number} | undefined>(undefined)

// On load, set from query param if available
watchEffect(() => {
  if (!genres.value.length) return
  const genreId = Number(route.query.genre)
  if (genreId) {
    const found = genres.value.find(g => g.id === genreId)
    if (found) selectedGenre.value = { label: found.genre_name, value: found.id }
    else selectedGenre.value = undefined
  } else {
    selectedGenre.value = undefined
  }
})

// Filter books by selected genre
const filteredBooks = computed(() => {
  if (!selectedGenre.value) return data
  return data?.filter(book =>
    (book.genres || []).some(g => g.id === Number(selectedGenre.value?.value))
  )
})

watch(selectedGenre, (val) => {
  if (val?.value) {
    router.replace({ query: { ...route.query, genre: val.value } })
  } else {
    const { genre, ...rest } = route.query
    router.replace({ query: rest })
  }
})

</script>

<template>
  <main class="flex flex-col gap-4 items-center my-10">
    <!-- Genre Filter -->
    <div class="mb-6">
      <UInputMenu
        v-model="selectedGenre"
        :items="genres.map((genre) => ({ label: genre.genre_name, value: genre.id }))"
        placeholder="Filter by genre"
        clearable
        :ui="{
          trailingIcon: 'hover:cursor-pointer',
          item: 'hover:cursor-pointer',
        }"
      />
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-6 gap-4">
      <UCard
        v-for="n in 6"
        :key="n"
        class="border border-gray-200 overflow-hidden flex flex-col w-56"
      >
        <USkeleton class="w-full h-48 mb-4" />
        <USkeleton class="h-6 w-3/4 mb-2" />
        <USkeleton class="h-4 w-1/2 mb-2" />
      </UCard>
    </div>

    <!-- Error State -->
    <p v-else-if="error">Something went wrong. Please try again later.</p>

    <!-- Data State -->
    <ul v-else-if="filteredBooks" class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-6 gap-4">
      <li v-for="book in filteredBooks" :key="book.id">
        <ULink :href="`/books/${book.slug}`">
          <UCard
            class="border border-gray-200 overflow-hidden flex flex-col hover:cursor-pointer"
            :ui="{
              header: 'p-0 sm:px-0',
              body: 'p-3'
            }"
          >
            <template #header>
              <nuxt-img
                :src="`http://localhost:3001/static/covers/${book.slug}.jpg`"
                class="w-full h-72 object-cover"
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
                  class="bg-indigo-100 text-indigo-800 rounded-full px-2 py-1 text-xs text-center font-semibold mt-2"
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