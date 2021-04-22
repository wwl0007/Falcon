<template>
    <div id="app">
        <Nav></Nav>
        <div class="container">
            <router-view />
        </div>
    </div>
</template>

<script lang="ts">
    import { Component, Vue } from 'vue-property-decorator'
    import Nav from '@/components/Nav/Nav.vue'

    @Component({
        components: {
            Nav
        }
    })
    export default class App extends Vue {
    }
</script>

<style>
    #app {
        font-family: Avenir, Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        color: #2c3e50;
    }
</style>
