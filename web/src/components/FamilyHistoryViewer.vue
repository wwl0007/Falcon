<template>
    <div class="d-flex flex-column mb-4">
        Family History
        <div>
            <b-button>{{ showButtonText }}</b-button>
        </div>

        <div v-show="isShown">
            <b-table :items="value">

            </b-table>
        </div>
    </div>
</template>

<script lang="ts">
    import { Component, Prop, Vue } from 'vue-property-decorator'
    import { RelativeHistoryItem } from '@/api/patients'

    @Component
    export default class FamilyHistoryViewer extends Vue {
        @Prop() value!: RelativeHistoryItem[];

        isShown = false;

        get showButtonText() {
            return this.isShown ? 'Hide' : 'Show';
        }
    }
</script>

<style scoped>

</style>
