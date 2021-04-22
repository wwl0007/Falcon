<template>
    <div class="d-inline-flex justify-content-between mt-3 pt-1 pb-1 my-class border-bottom">
        <span class="me-4 fw-bolder">{{ label }}:</span>
        <div class="d-inline-flex">
            <span :class="`ms-4 ${valueClass}`">{{ value ? value : '[ blank ]'}}</span>
            <i class="ms-2 bi-pencil text-muted" style="font-size: .8rem" />
        </div>
    </div>
</template>

<script lang="ts">
    import { Component, Prop, Vue } from 'vue-property-decorator'

    @Component
    export default class EditableField extends Vue {
        @Prop() label!: string;
        @Prop() value!: string;

        get valueClass() {
            return this.value ? '' : 'text-muted';
        }
    }
</script>

<style scoped>
    .my-class {
        /*border-bottom-width: 2px;*/
        /*border-bottom-style: dot-dash;*/
        /*border-bottom-color: #2c3e50;*/
    }
</style>
