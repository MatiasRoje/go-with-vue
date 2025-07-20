<script setup lang="ts">
import type { NavigationMenuItem } from '@nuxt/ui'
import { useAuth } from '~/composables/useAuth'
import { useCookie } from '#app'

const { isAuthenticated, user, loading } = useAuth()
console.log("Header.vue loading ref:", loading)
console.log("Header.vue loading value:", loading.value)

const handleLogout = () => {
  useAuth().user.value = null
  useAuth().token.value = null
  useCookie("user", {secure: true}).value = null
  useCookie("token", {secure: true}).value = null
}

// Main navigation group (SSR-friendly)
const mainItems: NavigationMenuItem[] = [
  {
    label: 'Home',
    icon: 'i-lucide-home',
    to: '/',
  },
  // Add more main items here
]

// User/auth group (client-only)
const userItems = computed<NavigationMenuItem[]>(() =>
  isAuthenticated.value
    ? [
        {
          label: user.value?.email ?? "",
          icon: "i-lucide-user",
          class: "pointer-events-none",
        },
        {
          label: "Logout",
          icon: "i-lucide-log-out",
          onClick: handleLogout,
          class: "hover:cursor-pointer"
        }
      ]
    : [
        {
          label: "Login",
          icon: "i-lucide-log-in",
          to: "/login",
        }
      ]
)
</script>

<template>
  <div class="flex items-center w-full border-b border-default p-2 justify-between">
    <!-- Main nav: always SSR -->
    <UNavigationMenu
      :items="[mainItems]"
      highlight
      highlight-color="primary"
      orientation="horizontal"
      class="flex-1"
    />

    <!-- User nav: client-only, with loading skeleton -->
    <ClientOnly>
      <!-- This content is rendered only on the client -->
      <div class="min-w-[120px] flex items-center">
        <div v-if="loading" class="flex items-center gap-2 mr-4">
          <USkeleton class="h-6 w-6 rounded-full" />
          <USkeleton class="w-24 h-6 rounded" />
        </div>
        <UNavigationMenu
          v-if="!loading"
          :items="[userItems]"
          highlight
          highlight-color="primary"
          orientation="horizontal"
        />
      </div>
      <!-- This is the fallback content, rendered on the server -->
      <template #fallback>
        <div class="flex items-center gap-2 mr-4">
          <USkeleton class="h-6 w-6 rounded-full" />
          <USkeleton class="w-24 h-6 rounded" />
        </div>
      </template>
    </ClientOnly>
  </div>
</template>