<template>
    <div>
        <div class="d-flex justify-content-end">
            <b-button variant="primary" @click="$router.push('/patients/edit')">Add Patient</b-button>
        </div>
        <b-table striped hover :items="items" :fields="fields">
            <template #cell(Action)="row">
                <b-button @click="$router.push(`/patients/view/${row.item.ID}`)">View</b-button>
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
            'Pathogenic',
            'Gene',
            'HistoryClass',
            'Ethnicity',
            'ConsentApproval',
            'CancerDX',
            'CancerDXType',
            'CancerDXAge',
            'KnownBRCA',
            'KnownCancer',
            'Action'
        ]

        async mounted() {
            this.items = await fetchPatients();
        }

        doThing(val: any) {
            console.log(val);
        }
    }
</script>

<style scoped>

</style>
