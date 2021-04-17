<template>
    <div class="d-flex flex-column">
        <div class="d-flex justify-content-center">
            <img alt="Vue logo" src="../assets/logo.png" width="128px" height="128px" class="mb-4">
        </div>
        <div class="flex-column justify-content-between">
            <h2>Welcome to our hackathon project.</h2>
            <div>
                Our group members:
                <ul>
                    <li>Wesley</li>
                    <li>Bill</li>
                    <li>Martin</li>
                    <li>Jordan Davis</li>
                    <li>Michael Porter</li>
                </ul>
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
