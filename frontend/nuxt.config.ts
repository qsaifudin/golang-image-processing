export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: ["vuetify-nuxt-module",],
  runtimeConfig: { public: { BASE_URL: process.env.BASE_URL } },
  vuetify: {
    moduleOptions: {},
    vuetifyOptions: {
      theme: {
        defaultTheme: "dark",
        themes: {},
      },
    },
  },
});
