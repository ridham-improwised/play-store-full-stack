<script setup>
const props = defineProps(['label', 'forId', 'placeholder', 'invalidField', 'modelValue'])
const emits = defineEmits(['update:modelValue'])
</script>

<template>
    <div>
        <label :for="forId" class="form-label mb-0">{{ label }}</label>
        <input type="text" class="form-control checkField" :id="forId" :placeholder="placeholder"
            @input="$emit('update:modelValue', $event.target.value)" :value="modelValue"
            required>
        <div class="invalid-feedback">
            {{ invalidField }}
        </div>
    </div>
</template>
