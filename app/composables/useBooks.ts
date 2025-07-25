import type { BackendResponse, Book } from "~/types/types"

export const useBooks = () => {
  const getBooks = async () => {
    const {data, pending, error, refresh} = await useFetch<BackendResponse<{books: Book[]}>>("/books", {
      baseURL: useRuntimeConfig().public.API_URL,
    })
    return {
      data: data.value?.data?.books,
      error,
      loading: pending.value,
      refresh
    }
  }

  const getBook = async (slug: string) => {
    const {data, pending, error, refresh} = await useFetch<BackendResponse<{book: Book}>>(`/books/${slug}`, {
      baseURL: useRuntimeConfig().public.API_URL,
    })
      return {
        data: data.value?.data?.book,
        error,
        loading: pending.value,
        refresh
      }
    }

  return {
    getBooks,
    getBook
  }
}
