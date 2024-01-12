// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  runtimeConfig: {
    public: {
      base_url: process.env.BASE_URL,
      random_img_access_key: process.env.RANDOM_IMAGE_ACCESS_KEY
    },
  },
  app: {
    head: {
      title: "Google Play Store",
      meta: [
        { charset: "utf-8" },
        { name: "viewport", content: "width=device-width, initial-scale=1" },
      ],
      link: [
        {rel:"stylesheet", href:"https://cdn.jsdelivr.net/npm/select2-bootstrap-5-theme@1.3.0/dist/select2-bootstrap-5-theme.min.css"}
      ]
    },
  },
  plugins: [
    '~/plugins/fontawesome.js',
    { src: '~/plugins/bootstrap.client.ts', mode: 'client' }
  ],
  css: [
    "bootstrap/dist/css/bootstrap.min.css",
    "~/assets/css/style.css",
    "select2/dist/css/select2.css",
  ],
  vite: {
    define: {
      "process.env.DEBUG": false,
    },
  },
});
