<script setup lang="ts">
import type { NavigationMenuItem } from '@nuxt/ui'

const items = ref<NavigationMenuItem[]>([
  {
    label: 'Home',
    icon: 'i-lucide-home',
    to: '/',
  },
  {
    label: 'Login',
    icon: 'i-lucide-user',
    to: '/login',
  }
])
</script>

<template>
  <UNavigationMenu :items="items" class="w-full justify-center" variant="pill"/>
</template>
