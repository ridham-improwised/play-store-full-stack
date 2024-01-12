<script setup>
const router = useRouter() ;
const { messages, messageRemover } = allApiMessages();
onMounted(() => {
    setTimeout(() => {
        if(messages.value.updateStatus) {
            messageRemover()
            router.back() ;
        }
        messageRemover()
    }, 3000)
})
</script>

<template>
    <div class="text-center">
        <!-- Error -->
        <div v-if="messages.apiError" ref="failure" class="mb-0 alert alert-danger alert-dismissible fade show"
            role="alert">
            <h2 class="m-0">Something went wrong !!!</h2>
            <p class="m-0">{{ messages.apiError }}</p>
        </div>

        <div v-else class="mb-0 alert alert-success alert-dismissible fade show"
            role="alert">
            <p class="m-0" v-if="messages.addStatus">{{ messages.addStatus }}</p>
            <p class="m-0" v-if="messages.updateStatus">{{ messages.updateStatus }}</p>
            <p class="m-0" v-if="messages.deleteStatus">{{ messages.deleteStatus }}</p>
            <p class="m-0" v-if="messages.reviewStatus">{{ messages.reviewStatus }}</p>
        </div>
    </div>
</template>
