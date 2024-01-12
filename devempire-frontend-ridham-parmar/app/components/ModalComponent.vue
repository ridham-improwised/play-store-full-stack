<script setup>
import { Modal } from 'bootstrap';
const { deleteApp, addReview } = callApi();
const { messages } = allApiMessages();
const props = defineProps(['deleteAppId', 'reviewModal', 'app']);
const isMarkedRating = ref(false)
const numRating = ref(0)
const reviewAlert = ref(false)
const isReviewModal = computed(() => {
    return props.reviewModal
})
const formData = reactive({
    app: "",
    translated_Review: "",
    sentiment: ""
})

async function handleReviewSubmit() {
    let modal ;
    if (process.client) {
        modal = Modal.getOrCreateInstance(document.getElementById('modalFeature'))
    }
    if (formData.translated_Review === "" || !numRating.value) {
        reviewAlert.value = true
        return
    }
    modal.hide();
    formData.app = props.app
    formData.sentiment = numRating.value >= 3 ? "Positive" : "Negative"
    await addReview(formData);
}
const vReviewAlert = {
    mounted: async (el) => {
        setTimeout(() => {
            reviewAlert.value = false
        }, 2000)
    }
}
</script>

<template>
    <ClientOnly>
        <div>
            <div class="modal fade" id="modalFeature" tabindex="-1" aria-labelledby="modalLabel" aria-hidden="true">
                <div class="modal-dialog modal-dialog-centered">
                    <!-- modal body for add review -->
                    <div class="modal-content" v-if="isReviewModal">
                        <div class="modal-header">
                            <h1 class="modal-title fs-5" id="exampleModalLabel">{{ app }}</h1>
                            <button type="button" class="btn-close" data-bs-dismiss="modal"
                                @click="isMarkedRating = false, numRating = 0" aria-label="Close"></button>
                        </div>
                        <div class="modal-body">
                            <form @submit.prevent="handleReviewSubmit">
                                <div class="mb-3 d-flex justify-content-evenly fs-4">
                                    <div v-for="n in 5" :key="n">
                                        <font-awesome-icon icon="star" class="modal-stars text-secondary"
                                            @click="isMarkedRating = true, numRating = n"
                                            :class="{ 'text-success': isMarkedRating && n <= numRating }" />
                                    </div>
                                </div>

                                <div class="mb-4">
                                    <textarea class="form-control" id="message-text" rows="5"
                                        placeholder="Describe your experience"
                                        v-model.trim="formData.translated_Review"></textarea>
                                </div>

                                <div class="alert alert-danger reviewAlert" v-if="reviewAlert" v-review-alert role="alert">
                                    Please fill required details !!
                                </div>

                                <div class="d-flex gap-3 justify-content-end">
                                    <button type="button" class="btn btn-secondary"
                                        @click="isMarkedRating = false, numRating = 0"
                                        data-bs-dismiss="modal">Close</button>
                                    <button type="submit" class="btn btn-primary">Send
                                        message</button>
                                </div>
                            </form>
                        </div>
                    </div>

                    <!-- modal body for delete app -->
                    <div class="modal-content" v-else>
                        <div class="modal-header">
                            <h1 class="modal-title fs-5" id="deleteModalLabel">Delete App</h1>
                            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                        </div>
                        <div class="modal-body">
                            Are you sure you want to delete app and its reviews ?
                        </div>
                        <div class="modal-footer">
                            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">No</button>
                            <button type="button" class="btn btn-primary" data-bs-dismiss="modal"
                                @click="deleteApp(deleteAppId)">Yes, Delete</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </ClientOnly>
</template>
