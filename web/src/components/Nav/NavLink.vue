<template>
    <router-link :to="to" :class="`nav-link px-2 ${linkType}`"><slot></slot></router-link>
</template>

<script lang="ts">
    import { Component, Prop, Vue, Watch } from 'vue-property-decorator'
    import { Route } from 'vue-router'

    @Component
    export default class NavLink extends Vue {
        @Prop() to!: string;

        currentRoutePath: string = this.$router.currentRoute.path;

        get linkType(): string {
            return this.currentRoutePath === this.to ? 'active' : '';
        }

        @Watch('$route')
        routeChanged(newRoute: Route): void {
            this.currentRoutePath = newRoute.path;
        }
    }
</script>

<style scoped>

</style>
