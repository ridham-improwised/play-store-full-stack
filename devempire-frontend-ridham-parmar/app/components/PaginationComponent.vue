<script setup>
const route = useRoute() ;
const props = defineProps(['page', 'pagination', 'navigationPath']);
console.log("pagination ", props.pagination);

function selectPageHandler(pageNo) {
    let navigateQueryData = {} ;
    if(route.query.type) {
        Object.assign(navigateQueryData, {type: route.query.type})
    }
    if(route.query.category) {
        Object.assign(navigateQueryData, {category: route.query.category})
    }
    if(route.query.search) {
        Object.assign(navigateQueryData, {search: route.query.search})
    }
    if(pageNo >= 1 && pageNo !== props.page) {
        navigateTo({ path: props.navigationPath, query: Object.assign({ page: pageNo }, navigateQueryData) })
    }
}
</script>

<template>
    <nav aria-label="Page navigation example">
        <ul class="pagination justify-content-center">
            <li class="page-item">
                <NuxtLink class="page-link" :class="{ disabled : page == 1}"  @click="selectPageHandler(page - 1)">Previous
                </NuxtLink>
            </li>

            <li class="page-item" v-for="n in 2" :key="n">
                <NuxtLink class="page-link" v-if="pagination.currentPage < pagination.totalPages"
                    :class="{ active: page + (n - 1) === page }" @click="selectPageHandler(page + n - 1)">
                    {{ page + n - 1 }}
                </NuxtLink>
            </li>

            <li class="page-item">
                <NuxtLink class="page-link" :class="{ disabled : pagination.currentPage >= pagination.totalPages}"  @click="selectPageHandler(page + 1)">Next
                </NuxtLink>
            </li>
        </ul>
    </nav>
</template>
