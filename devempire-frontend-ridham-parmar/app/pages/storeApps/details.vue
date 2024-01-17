<script setup>
const route = useRoute();
const { messages } = allApiMessages();
const { fetchAppDetails } = callApi();
const reviewsLength = ref(6)
const reviewModal = ref(false)

watch(route, () => {
    fetchAppDetails(parseInt(route.query.id), route.fullPath.split('=').at(2).split('.').join(' '))
}, { immediate: true })
onUnmounted(() => {
    messages.value.appDetails = null;
    messages.value.appReviews = null;
})
watch(() => [messages.value.reviewStatus], () => {
    if (messages.value.reviewStatus === null) {
        fetchAppDetails(parseInt(route.query.id), route.fullPath.split('=').at(2).split('.').join(' '))
    }
})
</script>

<template>
    <ClientOnly>
        <div class="row flex-column justify-content-center g-0 w-100">
            <div class="col-12 col-lg-10 d-flex flex-column align-self-center gap-5">
                <MessageDisplay v-if="messages.apiError || messages.reviewStatus" />
                <LoadingComponent v-else-if="messages.pendingStatus" />
                <div v-else-if="messages.appDetails && messages.appReviews">
                    <div class="d-flex flex-column gap-5">
                        <!-- top section-->
                        <div class="row g-0">
                            <div class="col-12 col-md-8 d-flex flex-column gap-4 justify-content-evenly">
                                <div class="d-flex gap-3">
                                    <div class="d-block d-md-none">
                                        <ImageDisplay :image="messages.appDetails.App" />
                                    </div>
                                    <div class="">
                                        <h1 class="mb-md-4 topHeading">{{ messages.appDetails.App }}</h1>
                                        <p class="text-secondary mb-0 detailsPara">{{
                                            messages.appDetails.Category.at(0).toUpperCase() +
                                            (messages.appDetails.Category).toLowerCase().split('_').join(" ").slice(1) }}
                                        </p>
                                    </div>
                                </div>

                                <div
                                    class="d-flex gap-2 gap-sm-3 justify-content-evenly justify-content-sm-start text-center">
                                    <div v-if="messages.appDetails.Rating || messages.appDetails.Reviews"> 
                                        <p class="top-text mb-0 detailsPara">
                                            {{ messages.appDetails.Rating }}
                                            <font-awesome-icon icon="star" />
                                        </p>
                                        <p class="mb-0 detailsPara">
                                            {{ messages.appDetails.Reviews }} reviews
                                        </p>
                                    </div>
                                    <div class="vr align-self-center verticalRow" 
                                        v-if="messages.appDetails.Rating || messages.appDetails.Reviews">
                                    </div>
                                    <div v-if="messages.appDetails.Installs">
                                        <p class="top-text mb-0 detailsPara">
                                            {{ messages.appDetails.Installs }}
                                        </p>
                                        <p class="mb-0 detailsPara">Downloads</p>
                                    </div>
                                    <div class="vr align-self-center verticalRow" 
                                        v-if="messages.appDetails.Installs">
                                    </div>
                                    <div>
                                        <p class="top-text mb-0 detailsPara">
                                            <font-awesome-icon icon="e"
                                                v-if="messages.appDetails['Content Rating'] == 'everyone'" />
                                            <font-awesome-icon icon="t" v-else />
                                        </p>
                                        <p class="mb-0 detailsPara">
                                            {{ messages.appDetails['Content Rating'].at(0).toUpperCase() +
                                                messages.appDetails['Content Rating'].slice(1) }}</p>
                                    </div>
                                    <div class="vr align-self-center verticalRow"></div>
                                    <div>
                                        <p class="top-text mb-0 detailsPara">{{ messages.appDetails.Size.split('M').at(0) }}
                                            MB</p>
                                        <p class="mb-0 detailsPara">Size</p>
                                    </div>
                                </div>
                            </div>
                            <div class="d-none col-md-4 d-md-flex justify-content-end align-items-center">
                                <ImageDisplay :image="messages.appDetails.App" />
                            </div>
                        </div>

                        <!-- add rating section -->
                        <div class="row gap-2 flex-column flex-md-row align-items-center g-0">
                            <div class="col">
                                <h4 class="mb-0">Rate this app</h4>
                                <p class="text-secondary mb-0 detailsPara">Tell others what you think</p>
                            </div>
                            <div>
                                <ModalComponent :reviewModal="reviewModal" :app="messages.appDetails.App" />
                                <p class="text-primary opacity-100 mb-0 fw-bold detailsPara" @click="reviewModal = true"
                                    data-bs-toggle="modal" data-bs-target="#modalFeature">Write a review</p>
                            </div>
                        </div>

                        <!-- reviews section -->
                        <div class="row gap-2 flex-column flex-md-row align-items-center g-0">
                            <div class="col">
                                <div class="row flex-column flex-md-row align-items-center g-0">
                                    <div class="col">
                                        <div
                                            class="d-flex justify-content-between justify-content-md-start align-items-center gap-4">
                                            <h4 class="mb-0">Rating and reviews</h4>
                                            <font-awesome-icon icon="arrow-right" />
                                        </div>
                                    </div>
                                    <div class="col">
                                        <div class="d-flex align-items-center justify-content-md-end gap-3">
                                            <p class="mb-0 detailsPara">Ratings and reviews are verified</p>
                                            <font-awesome-icon icon="circle-info" />
                                        </div>
                                    </div>
                                </div>
                            </div>

                            <template v-if="messages.appReviews.length > 0">
                                <ReviewComponent :reviews="messages.appReviews.slice(0, reviewsLength)" />

                                <h6 class="text-primary opacity-100 mb-0" v-if="messages.appReviews.length > 6"
                                    @click="reviewsLength = reviewsLength === 6 ? messages.appReviews.length : 6">
                                    {{ reviewsLength === 6 ? "See all reviews" : "See less" }}
                                </h6>
                            </template>

                            <div v-else>
                                <p class="mb-0 detailsPara text-secondary">No ratings and reviews exits</p>
                            </div>
                        </div>

                        <!-- about section -->
                        <div class="row gap-2 flex-column g-0">
                            <div class="col">
                                <div
                                    class="d-flex justify-content-between justify-content-md-start align-items-center gap-4">
                                    <h4 class="mb-0">About this app</h4>
                                    <font-awesome-icon icon="arrow-right" />
                                </div>
                            </div>
                            <div class="col-12 col-md-8">
                                <div class="row g-0 mb-4 gap-5">
                                    <div class="col">
                                        <p class="mb-0 detailsPara fw-normal">Version</p>
                                        <p class="mb-0 detailsPara text-secondary">{{ messages.appDetails['Current Ver'] }}
                                        </p>
                                    </div>
                                    <div class="col">
                                        <p class="mb-0 detailsPara fw-normal">Updated On</p>
                                        <p class="mb-0 detailsPara text-secondary">{{ messages.appDetails['Last Updated'] }}
                                        </p>
                                    </div>
                                </div>
                                <div class="row g-0 mb-4 gap-5">
                                    <div class="col">
                                        <p class="mb-0 detailsPara fw-normal">Requires android</p>
                                        <p class="mb-0 detailsPara text-secondary">{{ messages.appDetails['Android Ver'] }}
                                            and up</p>
                                    </div>
                                    <div class="col">
                                        <p class="mb-0 detailsPara fw-normal">{{ messages.appDetails.Type }}</p>
                                        <p class="mb-0 detailsPara text-secondary" v-if="messages.appDetails.Price">{{
                                            messages.appDetails.Price }}</p>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </ClientOnly>
</template>
