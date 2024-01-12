<script setup>
const filterData = inject('filterData')
</script>

<template>
    <div>
        <input class="form-control me-2" v-model="filterData.search"
            placeholder="Search apps e.g. Instagram" aria-label="Search" />
    </div>
</template>
