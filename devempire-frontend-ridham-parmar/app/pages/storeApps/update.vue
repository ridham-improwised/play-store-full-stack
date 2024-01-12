<script setup>
const { messages } = allApiMessages();
</script>

<template>
    <div class="w-100">
        <MessageDisplay v-if="messages.updateStatus || messages.apiError" />
        <LoadingComponent v-else-if="messages.pendingStatus" />
        <FormComponent v-else/>
    </div>
</template>
