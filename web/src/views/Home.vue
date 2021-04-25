<template>
    <div class="d-flex justify-content-center">
        <div class="d-flex flex-column">
            <div class="d-flex justify-content-center">
                <img alt="Falcon logo" src="../assets/logo_bird_with_name.png" width="240px" height="240px" class="mb-4">
            </div>
            <div class="flex-column justify-content-between">
                <h1>You have reached <span class="fw-bold">Project Falcon</span>.</h1>
                <div class="mt-4">
                    <p>
                        Our group members:
                    </p>
                    <ul>
                        <li>Bill Luu</li>
                        <li>Wesley Lowman</li>
                        <li>Jordan Davis</li>
                        <li>Michael Porter</li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
    import { Component, Vue } from 'vue-property-decorator'

    @Component
    export default class Home extends Vue {
    }
</script>
