<template>
    <div class="card mt-2 mb-2 p-2 shadow-sm">
        <div class="d-inline-flex justify-content-between">
            <LineItem class="fs-5 text-black-50 fw-bold" v-model="value.Relation" @input="onValueChange" :edit-mode="editMode" />
            <div v-if="!editMode">
                <span class="link-primary">
                    <a class="text-decoration-none" href="javascript:void(0)" @click="editMode=true">
                        Edit
                    </a>
                </span>
            </div>
        </div>
        <LineItem v-model="value.Cancer" @input="onValueChange" :edit-mode="editMode" />
        <LineItem v-model="value.Age" append=" years old" @input="onValueChange" :edit-mode="editMode" />
        <div v-if="editMode" class="d-inline-flex justify-content-between pt-2">
            <div>
                <b-button class="me-2" size="sm" variant="danger" @click="editMode=false">Cancel</b-button>
                <b-button size="sm">Save</b-button>
            </div>
            <div>
                <b-button size="sm" variant="danger">Delete</b-button>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
    import { Component, Prop, Vue } from 'vue-property-decorator'
    import { RelativeHistoryItem } from '@/api/patients'
    import LineItem from '@/components/FamilyHistory/LineItem.vue'

    @Component({
        components: { LineItem }
    })
    export default class FamilyHistoryItem extends Vue {
        @Prop() value!: RelativeHistoryItem;

        editMode = false;

        onValueChange() {

        }
    }
</script>

<style scoped>

</style>
