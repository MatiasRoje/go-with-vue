import { useAuth } from "~/composables/useAuth";
import { useCookie } from "#app";
import type { BackendResponse, User } from "~/types/types";

export default defineNuxtPlugin(() => {
  const auth = useAuth();
  const token = useCookie<string | null>("token", {secure: true});
  const userCookie = useCookie<string | null>("user", {secure: true});

  auth.loading.value = true;

  const validateToken = async () => {
    // Only try to validate if token exists
    if (token.value) {
      try {
        const response = await $fetch("/validate-token", {
          baseURL: useRuntimeConfig().public.API_URL,
          method: "POST",
          headers: { Authorization: `Bearer ${token.value}` },
        }) as BackendResponse<{user: User}>;
        if (!response.error && response.data) {
          // If valid, set user again in global state as we always want to have the latest user data
          auth.user.value = response.data.user;
          auth.token.value = token.value;
          userCookie.value = JSON.stringify(response.data.user)
          await new Promise((resolve) => setTimeout(resolve, 1000));
        } else {
          // Invalid token, clear everything
          auth.user.value = null;
          auth.token.value = null;
          token.value = null;
          userCookie.value = null;
        }
      } catch (error) {
        // If error, clear everything and log error
        console.error("Token validation error:", error);
        auth.user.value = null;
        auth.token.value = null;
        token.value = null;
        userCookie.value = null;
      }
    }
    auth.loading.value = false;
  }

  validateToken()
});
