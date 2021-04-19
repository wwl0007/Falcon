<template>
    <div>
        <div class="d-flex justify-content-end">
            <b-button variant="primary" @click="$router.push('/patients/edit')">Add Patient</b-button>
        </div>
        <b-table striped hover :items="items" :fields="fields">
            <template #cell(RelativeHistory)="">
                <b-button>View</b-button>
            </template>
        </b-table>
    </div>
</template>

<script lang="ts">
    import { Component, Vue } from 'vue-property-decorator'
    import { fetchPatients, Patient } from '@/api/patients.ts'

    @Component
    export default class Patients extends Vue {
        items: Patient[] = [];
        fields = [
            'ID',
            'IsPathogenic',
            'Gene',
            'HistoryClass',
            'ethnicity',
            'consent_approval',
            'cancer_dx',
            'cancer_dx_type',
            'cancer_dx_age',
            'known_brca',
            'known_cancer',
            'RelativeHistory'
        ]

        async mounted() {
            this.items = await fetchPatients();
        }
    }
</script>

<style scoped>

</style>
