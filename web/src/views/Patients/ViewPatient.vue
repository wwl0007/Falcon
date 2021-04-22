<template>
    <div class="border rounded bg-body p-4">
        <div class="mb-4">
            <b-button @click="$router.back()">Back</b-button>
        </div>
        <div class="d-flex justify-content-center">
            <div v-if="patient !== null" style="min-width: 50%">
                <div class="d-flex flex-column">
                    <div class="d-inline-flex justify-content-center mb-2">
                        <span class="h2 me-4">Patient ID #{{ patientId }}</span>
                    </div>
                    <EditableField label="Pathogenic" :value="patient.Pathogenic" />
                    <EditableField label="Gene" :value="patient.Gene" />
                    <EditableField label="History Class" :value="patient.HistoryClass" />
                    <EditableField label="Ethnicity" :value="patient.Pathogenic" />
                    <EditableField label="Consent Approval" :value="patient.ConsentApproval" />
                    <EditableField label="Cancer DX" :value="patient.CancerDX" />
                    <EditableField label="Cancer DX Type" :value="patient.CancerDXType" />
                    <EditableField label="Cancer DX Age" :value="patient.CancerDXAge" />
                    <EditableField label="Known BRCA" :value="patient.KnownBRCA" />
                    <EditableField label="Known Cancer" :value="patient.KnownCancer" />
                </div>
                <div class="mt-4 mb-4">
                    <p class="h3">Family History</p>
                    <FamilyHistoryItem
                        v-for="familyMember of patient.RelativeHistory"
                        :key="familyMember.ID"
                        :value="familyMember"
                    />
                </div>
            </div>
            <div v-else>
                Loading patient...
            </div>
        </div>
    </div>
</template>

<script lang="ts">
    import { Component, Vue } from 'vue-property-decorator'
    import { fetchPatientById, Patient } from '@/api/patients'
    import EditableField from '@/components/EditableField.vue'
    import FamilyHistoryViewer from '@/components/FamilyHistory/FamilyHistoryViewer.vue'
    import FamilyHistoryItem from '@/components/FamilyHistory/FamilyHistoryItem.vue'
    @Component({
        components: { FamilyHistoryItem, FamilyHistoryViewer, EditableField }
    })
    export default class ViewPatient extends Vue {
        patient: Patient | null = null;

        get patientId() {
            return parseInt(this.$route.params.patientId);
        }

        async mounted() {
            this.patient = await fetchPatientById(this.patientId);
        }
    }
</script>

<style scoped>

</style>
