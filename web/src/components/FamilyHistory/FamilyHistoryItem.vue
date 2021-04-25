<template>
    <div class="card mt-2 mb-2 p-2 shadow-sm">
        <div class="d-inline-flex justify-content-between">
            <LineItem class="fs-5 text-black-50 fw-bold" v-model="value.Relation" :edit-mode="editMode" />
            <div v-if="!editMode">
                <span class="link-primary">
                    <a class="text-decoration-none" href="javascript:void(0)" @click="editMode=true">
                        Edit
                    </a>
                </span>
            </div>
        </div>
        <LineItem v-model="value.Cancer" :edit-mode="editMode" />
        <LineItem v-model="value.Age" type="number" append=" years old" :edit-mode="editMode" />
        <div v-if="editMode" class="d-inline-flex justify-content-between pt-2">
            <div>
                <b-button class="me-2" size="sm" variant="danger" @click="editMode=false">Cancel</b-button>
                <b-button size="sm" @click="onSave">Save</b-button>
            </div>
            <div>
                <b-button size="sm" variant="danger" @click="deleteItem">Delete</b-button>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
    import { Component, Prop, Vue } from 'vue-property-decorator'
    import { deleteRelativeHistory, RelativeHistoryItem, updateRelativeHistory } from '@/api/patients'
    import LineItem from '@/components/FamilyHistory/LineItem.vue'

    @Component({
        components: { LineItem }
    })
    export default class FamilyHistoryItem extends Vue {
        @Prop() value!: RelativeHistoryItem;

        editMode = false;

        async deleteItem() {
            await deleteRelativeHistory(this.value.ID);
            this.$emit('deleted', this.value);
        }

        async onSave() {
            try {
                await updateRelativeHistory(this.value);
                this.editMode = false;
            } catch {
                alert("Failed to save relative history item.");
            }
        }
    }
</script>

<style scoped>

</style>
