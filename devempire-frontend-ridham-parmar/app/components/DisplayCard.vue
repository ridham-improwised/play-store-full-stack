<script setup>
const props = defineProps(['appsData'])
const deleteAppId = ref()
const { getApps, fetchAppsImage } = callApi();
const vImageTag = {
    // mounted: async (el) => {
    //     el.src = await fetchAppsImage(el.alt)
    // }
}
</script>

<template>
    <div class="d-flex flex-column gap-4">
        
        <template v-if="appsData.data.apps">
            <NuxtLink :to="`/storeApps/details?id=${apps.id}&appname=${apps.app.split(' ').join('.')}`" @click.prevent v-for="(apps,index) in appsData.data.apps" :key="index" class="nav-link ">
                <div class="card">
                    <div class="row g-0 gap-lg-5">
                        <div class="col-md-4">
                            <img src="" :alt="apps.app" v-image-tag class="randomImg img-fluid rounded-end rounded-start h-100">
                        </div>
                        <div class="col-md-8 col-lg-6">
                            <div class="card-body d-flex flex-column justify-content-center h-100 gap-1">
                                <h5 class="card-title">{{ apps.app }}</h5>
                                <p class="card-text mb-2">
                                    {{ (apps.category).charAt(0).toUpperCase() + (apps.category).toLowerCase().split('_').join(" ").slice(1) }}
                                </p>
                                <div class="d-flex gap-3 mb-2">
                                    <span class="text-muted" v-if="apps.rating">
                                        {{ apps.rating}} 
                                        <font-awesome-icon icon="star" />
                                    </span>
                                    <span class="text-muted">{{ (apps.size).substring(0, (apps.size).length - 1) }} MB</span>
                                    <span class="text-muted" v-if="apps.type == 'Paid' ">
                                        {{ apps.price.charAt(0) === '$' ? apps.price : '$' + apps.price }}
                                    </span>
                                </div>
                                <div class="d-flex gap-3">
                                    <NuxtLink :to="`/storeApps/update?app=${apps.app.split(' ').join('.')}`">
                                        <button type="button" class="btn btn-success" @click="getApps(null, null, apps.id)">
                                            Update
                                        </button>
                                    </NuxtLink>
                                    <button type="button" @click.prevent @click="deleteAppId = apps.id" class="btn btn-danger"
                                        data-bs-toggle="modal" data-bs-target="#modalFeature">
                                        Delete
                                    </button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </NuxtLink>
            <ModalComponent :deleteAppId="deleteAppId" />
        </template>
    </div>
</template>
