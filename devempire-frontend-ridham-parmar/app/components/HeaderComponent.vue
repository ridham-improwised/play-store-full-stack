<script setup>
const route = useRoute();
const isBtnVisible = computed(() => {
    const path = route.path.split('/')
    return path.includes('storeApps') && path.length <= 2
})
</script>

<template>
    <nav class="navbar container-sm navbar-expand-md navbar-dark bg-body navPadding p-sm-0">
        <div class="container-fluid">
            <NuxtLink to="/">
                <img class="img-fluid playStoreLogo" src="../assets/logo/Google-Play-Logo.png" alt="Google-Play-Logo.png">
            </NuxtLink>
            <button class="navbar-toggler bg-secondary fs-6 px-2" type="button" data-bs-toggle="collapse"
                data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false"
                aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse justify-content-end navbar-collapse" id="navbarSupportedContent">
                <ul class="navbar-nav mb-lg-0 px-md-3 gap-lg-2">
                    <li class="nav-item">
                        <NuxtLink to="/" class="nav-link text-secondary" activeClass="active" aria-current="page">Home
                        </NuxtLink>
                    </li>
                    <li class="nav-item">
                        <NuxtLink to="/storeApps" class="nav-link text-secondary"
                            :class="{ active: $route.name.includes('storeApps') }" activeClass="active" aria-current="page">
                            Store Apps</NuxtLink>
                    </li>
                </ul>
                <NuxtLink to="/storeApps/add" v-if="isBtnVisible">
                    <ButtonComponent>
                        Add App
                    </ButtonComponent>
                </NuxtLink>
            </div>
        </div>
    </nav>
</template>
