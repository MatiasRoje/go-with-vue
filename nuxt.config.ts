// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2025-07-15",
  devtools: { enabled: true },
  modules: ["@nuxt/ui", "@nuxt/image", "@nuxt/eslint"],
  css: ["@/assets/css/main.css"],
  colorMode: {
    preference: "light",
    fallback: "light",
  },
  runtimeConfig: {
    public: {
      API_URL: process.env.API_URL,
    },
  },
});
