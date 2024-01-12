<script setup>
const props = defineProps(['value'])

const temp = computed(() => {
    return props.value
})
</script>

<template>
    <div class="d-flex gap-1">
        <div v-for="n in 5" :key="n">
            <font-awesome-icon v-if="n <= temp" icon="star" class="details-star text-success"/>
            <font-awesome-icon v-else icon="star" class="details-star text-secondary"/>
        </div>
    </div>
</template>

