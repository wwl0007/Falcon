<template>
    <header class="p-3 mb-3 border-bottom shadow-sm bg-body">
        <div class="container">
            <div class="d-flex flex-wrap align-items-center justify-content-center justify-content-lg-start">
<!--                <img-->
<!--                    alt="Falcon logo"-->
<!--                    src="@/assets/logo_bird_with_name.png"-->
<!--                    width="40px"-->
<!--                    height="40px"-->
<!--                    class="me-2 rounded-2 shadow-sm"-->
<!--                />-->
                <ul class="nav nav-pills col-12 col-lg-auto me-lg-auto mb-2 justify-content-center mb-md-0">
                    <li><NavLink to="/">Home</NavLink></li>
                    <li><NavLink to="/about">About</NavLink></li>
                    <li><NavLink to="/patients">Patients</NavLink></li>
                </ul>
            </div>
        </div>
    </header>
</template>

<script lang="ts">
    import { Component, Vue } from 'vue-property-decorator'
    import NavLink from '@/components/Nav/NavLink.vue'

    @Component({
        components: {
            NavLink
        }
    })
    export default class Nav extends Vue {
    }
</script>

<style scoped>

</style>
