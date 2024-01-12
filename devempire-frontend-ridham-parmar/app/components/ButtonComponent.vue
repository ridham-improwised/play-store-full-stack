<script setup>
const props = defineProps(['type'])
</script>

<template>
    <button class="w-100 btn btn-secondary" :type=" type ? 'submit' : null">
        <slot></slot>
    </button>
</template>
