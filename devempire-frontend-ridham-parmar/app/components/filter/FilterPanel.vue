<script setup>
const route = useRoute();
const props = defineProps(['navigationPath']);


const filterData = inject('filterData')
provide('filterData', filterData)

watch(route, () => {
    if (!route.query.search) {
        filterData.search = ''
    }
})

function searchHandle() {
    let querydata = {};
    if (filterData.search) {
        Object.assign(querydata, { search: filterData.search });
    }
    if (filterData.category) {
        Object.assign(querydata, { category: filterData.category });
    }
    if (filterData.type) {
        Object.assign(querydata, { type: filterData.type });
    }
    navigateTo({ path: props.navigationPath, query: querydata })
}
</script>

<template>
    <div class="mb-5">
        <form @submit.prevent="searchHandle" class="row mx-0 gap-3 gap-lg-4">

            <div class="col px-0">
                <div class="row">
                   <div class="col">
                    <FilterTypes :searchHandle="searchHandle" />
                   </div>
                    <div class="col">
                        <FilterCategories :searchHandle="searchHandle"/>
                    </div>
                </div>
            </div>
            <div class="col-lg-7 px-0">
                <div class="row justify-content-between">
                    <div class="col">
                        <FilterSearchComponent />
                    </div>
                    <div class="col-4 col-lg-3">
                        <ButtonComponent :type="true">Search</ButtonComponent>
                    </div>
                </div>
            </div>
        </form>
    </div>
</template>
