<script setup lang="ts">
import * as z from "zod"
import type { FormSubmitEvent } from "@nuxt/ui"
import type { BackendResponse, User } from "~/types/types"

const errorMessage = ref<string | null>(null)

const schema = z.object({
  email: z.string().email("Invalid email"),
  password: z.string().min(4, "Must be at least 4 characters")
})

type Schema = z.output<typeof schema>

const state = reactive<Partial<Schema>>({
  email: undefined,
  password: undefined
})

async function onSubmit(event: FormSubmitEvent<Schema>) {
  errorMessage.value = null
  
  const payload = {
    email: event.data.email,
    password: event.data.password
  }

  try {
    const response = await $fetch(`${useRuntimeConfig().public.API_URL}/login`, {
      method: "POST",
      body: payload
    }) as BackendResponse<{user: User, token: string}>
    if (response.data) {
      useAuth().user.value = response.data.user
      useAuth().token.value = response.data.token
      useCookie("user", {secure: true}).value = JSON.stringify(response.data.user)
      useCookie("token", {secure: true}).value = response.data.token
      navigateTo("/")
    }
  } catch (error) {
    console.error("Login error:", error)
    const backendErrorMessage = helpers.getBackendErrorMessage(error)
    if (backendErrorMessage) {
      errorMessage.value = backendErrorMessage
    }
  }
}

onUnmounted(() => {
  errorMessage.value = null
})
</script>

<template>
  <UForm :schema="schema" :state="state" class="space-y-4 w-72" @submit="onSubmit">
    <UFormField label="Email" name="email">
      <UInput v-model="state.email" class="w-full"/>
    </UFormField>

    <UFormField label="Password" name="password">
      <UInput v-model="state.password" type="password" class="w-full"/>
    </UFormField>

    <p v-if="errorMessage" class="text-red-500">{{ errorMessage }}</p>

    <UButton type="submit" class="hover:cursor-pointer">
      Submit
    </UButton>
  </UForm>
</template>
