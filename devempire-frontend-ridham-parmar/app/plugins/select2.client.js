import select2 from "select2";
export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.use(select2);
  return {
    provide: {
      select2Style: () => {
        $(".select2").select2({
          theme: "bootstrap-5",
        });
      },
    },
  };
});
