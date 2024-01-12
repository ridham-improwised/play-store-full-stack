<script setup>
const isScrolable = ref(false)
onMounted(() => {
    window.addEventListener('scroll', handleScroll)
})
const handleScroll = () => {
    isScrolable.value = window.scrollY > 0
}
</script>
<template>
    <div class="container-fluid d-flex align-items-center min-vh-100 position-relative">
        <div class="d-flex header navZIndex position-fixed start-0 end-0 top-0" :class="{'header-shadow': isScrolable}">
            <HeaderComponent />
        </div>
        <div class="container d-flex align-items-center justify-content-center nuxtPageContainer">
            <NuxtPage />
        </div>
        <div class="d-flex bg-dark position-absolute start-0 end-0 bottom-0">
            <FooterComponent/>
        </div>
    </div>
</template>

