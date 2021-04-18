<template>
    <div class="w-50">
        <div>
            <b-button @click="$router.back()">Back</b-button>
        </div>
        <h3 class="mt-2 text-center">Add new patient:</h3>
    </div>
</template>

<script lang="ts">
    import { Component, Vue } from 'vue-property-decorator'

    @Component
    export default class AddPatient extends Vue {

    }
</script>

<style scoped>

</style>
