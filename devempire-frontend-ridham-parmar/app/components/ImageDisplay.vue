<script setup>
const { fetchAppsImage } = callApi()
const props = defineProps(['image'])
const url = ref("../_nuxt/assets/image/loader.gif")
// onMounted(async () => {
//     url.value = await fetchAppsImage(props.image)
// })
</script>

<template>
    <img :src="`${url}`" :alt="image" class="randomImg appDetailImg">
</template>
