<script setup>
const { addApp, updateApp } = callApi();
const { messages } = allApiMessages();
const { validation } = validateForm();

const formData = reactive({
    app: "",
    category: "",
    size: "",
    type: "",
    price: "",
    contentRating: "",
    generes: "",
    currentVer: "",
    androidVer: ""
});

onMounted(() => {
    validation();
})

async function handleSubmit() {
    const fields = document.querySelectorAll('.checkField');
    let res = false;
    function checkInput() {
        res = (Array.from(fields)).every(checkInput);
        function checkInput(field) {
            return field.value.length > 0
        }
        return res;
    }
    if (checkInput()) {
        messages.value.appDetails ? await updateApp(formData, messages.value.appDetails.data.id) : await addApp(formData) ;
    }
}

watch(() => [messages.value.appDetails], () => {
    if (messages.value.appDetails) {
        ({
            app: formData.app, category: formData.category, size: formData.size, type: formData.type, price: formData.price,
            contentRating: formData.contentRating, generes: formData.generes,
            currentVer: formData.currentVer, androidVer: formData.androidVer
        } = messages.value.appDetails.data)
    }
}, {immediate: true})
</script>

<template>
    <div>
        <form class="row justify-content-center needs-validation " @submit.prevent="handleSubmit" novalidate>
            <div class="col-md-6 d-flex flex-column gap-3">
                <FormField label="App" forId="validationCustom01" v-model.trim="formData.app" placeholder="Instagram"
                    invalidField="Please provide app name." />

                <FormField label="Category" forId="validationCustom02" v-model.trim="formData.category"
                    placeholder="Social_Media" invalidField="Please provide category." />

                <FormField label="Size" forId="validationCustom03" v-model.trim="formData.size" placeholder="20 MB"
                    invalidField="Please provide size in MB." />

                <FormField label="Type" forId="validationCustom04" v-model.trim="formData.type" placeholder="Free/Paid"
                    invalidField="Please provide type." />

                <FormField label="Price" forId="validationCustom05" v-model.trim="formData.price" placeholder="$ 4.99"
                    invalidField="Please provide price." />

                <FormField label="Content Rating" forId="validationCustom06" v-model.trim="formData.contentRating"
                    placeholder="Everyone/Teen" invalidField="Please provide content rating." />

                <FormField label="Genres" forId="validationCustom07" v-model.trim="formData.generes"
                    placeholder="Social Media" invalidField="Please provide genres." />

                <FormField label="Current Version" forId="validationCustom08" v-model.trim="formData.currentVer"
                    placeholder="2.0.0" invalidField="Please provide current version." />

                <FormField label="Android Version" forId="validationCustom09" v-model.trim="formData.androidVer"
                    placeholder="Supported android version" invalidField="Please provide android version." />

                <div class="pt-3 d-flex justify-content-end">
                    <ButtonComponent :type="true">
                        Submit Form
                    </ButtonComponent>
                </div>
            </div>
        </form>
    </div>
</template>
