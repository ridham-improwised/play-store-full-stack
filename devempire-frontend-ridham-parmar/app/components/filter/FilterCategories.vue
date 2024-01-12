<script setup>
const { getCategory } = callApi();
const { messages } = allApiMessages();
const props = defineProps(['searchHandle']);
const filterData = inject('filterData')
const { $select2Style } = useNuxtApp();

onMounted(() => {
    $select2Style();
    $('.select-category').on('change', (event) => {
        filterData.category = event.target.value
        props.searchHandle();
    })
    $('.select2-container--bootstrap-5').css('width', '100%')
});

(function () {
    getCategory();
})();
</script>   

<template>
    <div >
        <select class="form-select select2 select-category" id="basic-usage">
            <option disabled selected>Categories</option>
            <option v-for="(category, index) in messages.categories" :value="category" :key="index">
                {{ category }}
            </option>
        </select>
    </div>
</template>
