<script setup lang="ts">
import * as z from 'zod'
import type { FormSubmitEvent } from '@nuxt/ui'

const schema = z.object({
  email: z.string().email('Invalid email'),
  password: z.string().min(4, 'Must be at least 4 characters')
})

type Schema = z.output<typeof schema>

const state = reactive<Partial<Schema>>({
  email: undefined,
  password: undefined
})

async function onSubmit(event: FormSubmitEvent<Schema>) {
  const payload = {
    email: event.data.email,
    password: event.data.password
  }

  const response = await $fetch(`${useRuntimeConfig().public.API_URL}/login`, {
    method: 'POST',
    body: payload
  })

  console.log(response)
}

</script>

<template>
  <UForm :schema="schema" :state="state" class="space-y-4 w-72" @submit="onSubmit">
    <UFormField label="Email" name="email">
      <UInput v-model="state.email" class="w-full"/>
    </UFormField>

    <UFormField label="Password" name="password">
      <UInput v-model="state.password" type="password" class="w-full"/>
    </UFormField>

    <UButton type="submit" class="hover:cursor-pointer">
      Submit
    </UButton>
  </UForm>
</template>
