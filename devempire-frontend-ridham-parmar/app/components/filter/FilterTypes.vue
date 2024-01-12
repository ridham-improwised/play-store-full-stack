<script setup>
const props = defineProps(['searchHandle']);
const filterData = inject('filterData')

onMounted(() => {
    $('.types').on('change', function (event) {
        filterData.type = event.target.value;
        props.searchHandle();
    });
})
</script>

<template>
    <select name="types" id="types" class="form-select types">
        <option disabled selected>Types</option>
        <option value="free">Free</option>
        <option value="paid">Paid</option>
    </select>
</template>
