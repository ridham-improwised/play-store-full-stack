<script setup>
const route = useRoute();
const { getApps } = callApi();
const { messages } = allApiMessages();
const page = ref();
const navigationPath = '/storeApps';

const filterData = reactive({
    search: '',
    type: '',
    category: ''
})
provide('filterData', filterData)

watch(route, () => {
    fetchApps()
}, { immediate: true })

watch(() => [messages.value.deleteStatus], () => {
    fetchApps()
})

function fetchApps() {
    let options = {};
    page.value = route.query.page === undefined ? 1 : parseInt(route.query.page);
    if (route.query.type) {
        Object.assign(options, { type: route.query.type })
    }
    if (route.query.category) {
        Object.assign(options, { category: route.query.category })
    }
    if (route.query.search) {
        Object.assign(options, { search: route.query.search })
    }
    getApps(page.value, options);
}
</script>

<template>
    <ClientOnly>
        <div class="row flex-column justify-content-center g-0">
            <div class="col-12 col-lg-10 d-flex flex-column mx-auto">
                <MessageDisplay v-if="messages.apiError || messages.deleteStatus" />
                <LoadingComponent v-else-if="messages.pendingStatus" />
                <div v-else-if="messages.success">
                    <FilterPanel :navigationPath="navigationPath" />
                    <DisplayCard :appsData="messages.success" />
                    <PaginationComponent :page="page" :pagination="messages.success.data.pagination" :navigationPath="navigationPath" />
                </div>
            </div>
        </div>
    </ClientOnly>
</template>
