<script setup>
const props = defineProps(['reviews'])
</script>

<template>
    <div class="row g-0 flex-column gap-2">
        <template v-for="(review, index) in reviews">
            <div v-if="review.Translated_Review && review.Translated_Review !== 'nan'" :key="index" 
                class="col-md-7 py-2 border-bottom border-secondary border-opacity-10">
                <template v-if="review.Sentiment === 'Positive'">
                    <StarComponent :value="Math.floor((Math.random() * 3) + 3)" />
                </template>
                <template v-if="review.Sentiment === 'Negative'">
                    <StarComponent :value="Math.floor(Math.random() * 2) + 1" />
                </template>
                <template v-if="review.Sentiment === 'Neutral'">
                    <StarComponent :value="3" />
                </template>
                <p class="text-secondary mb-0 detailsPara mt-1">{{ review.Translated_Review }}</p>
            </div>
        </template>
    </div>
</template>
