<template>
    <div class="d-inline-flex">
        <div v-if="!editing">
            <span>{{ value }}{{ append }}</span>
        </div>
        <div v-else-if="editMode" class="d-inline-flex">
            <b-input v-model="editValue" class="me-2" @blur="stopEdit" @keyup.enter="stopEdit" @keyup.tab="stopEdit" />
        </div>
        <div v-if="editMode && !editing">
            <i class="ms-2 bi-pencil text-muted" style="font-size: .8rem" @click="startEdit" />
        </div>
    </div>
</template>

<script lang="ts">
    import { Component, Prop, Vue, Watch } from 'vue-property-decorator'

    @Component
    export default class LineItem extends Vue {
        @Prop() value!: string;
        @Prop({ default: '' }) append!: string;
        @Prop({ default: false }) editMode!: boolean;

        editing = false;
        editValue: string | null = null;

        startEdit() {
            this.editValue = this.value;
            this.editing = true;
        }

        stopEdit() {
            this.editing = false;
            this.$emit('input', this.editValue);
        }

        @Watch('editMode')
        editModeChange(newValue: boolean) {
            if (!newValue) {
                this.editing = false;
            }
        }
    }
</script>

<style scoped>

</style>
