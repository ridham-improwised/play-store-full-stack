<script setup>
const { messages, setEmptyData } = allApiMessages();
onMounted(() => {
    setEmptyData(null)
})
</script>

<template>
    <div class="w-100">
        <MessageDisplay v-if="messages.addStatus || messages.apiError" />
        <LoadingComponent v-else-if="messages.pendingStatus"/>
        <FormComponent v-else/>
    </div>
</template>
