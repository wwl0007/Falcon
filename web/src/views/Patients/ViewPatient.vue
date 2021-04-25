<template>
    <div class="border rounded bg-body p-4">
        <div class="mb-4">
            <b-button @click="$router.back()">Back</b-button>
        </div>
        <div class="d-flex justify-content-center">
            <div v-if="patient !== null" style="min-width: 50%">
                <div class="d-flex flex-column">
                    <div class="d-inline-flex justify-content-center mb-2">
                        <span class="h2 me-4">Patient #{{ patientId }}</span>
                    </div>

                    <AICategorization
                        :confidence="93"
                        :value="riskFactor"
                        @input="riskFactorChanged"
                    />

                    <EditableField label="Pathogenic" :value="patient.Pathogenic" />
                    <EditableField label="Gene" :value="patient.Gene" />
<!--                    <EditableField label="History Class" :value="patient.HistoryClass" />-->
                    <EditableField label="Ethnicity" :value="patient.Ethnicity" />
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
                        @deleted="onHistoryItemDeleted"
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
    import { fetchPatientById, Patient, RelativeHistoryItem, updateRelativeHistory } from '@/api/patients'
    import EditableField from '@/components/EditableField.vue'
    import FamilyHistoryViewer from '@/components/FamilyHistory/FamilyHistoryViewer.vue'
    import FamilyHistoryItem from '@/components/FamilyHistory/FamilyHistoryItem.vue'
    import AICategorization from "@/components/AICategorization.vue";
    @Component({
        components: { AICategorization, FamilyHistoryItem, FamilyHistoryViewer, EditableField }
    })
    export default class ViewPatient extends Vue {
        patient: Patient | null = null;

        get patientId() {
            return parseInt(this.$route.params.patientId);
        }

        get riskFactor() {
            switch (this.patient?.HistoryClass) {
                case 'strong_personal':
                    return 100;
                case 'strong_family':
                    return 66;
                case 'not_strong':
                    return 33;
                default:
                    return 0;
            }
        }

        async mounted() {
            this.patient = await fetchPatientById(this.patientId);
        }

        onHistoryItemDeleted(item: RelativeHistoryItem) {
            if (this.patient) {
                this.patient.RelativeHistory = this.patient.RelativeHistory.filter(i => i.ID !== item.ID);
            }
        }

        riskFactorChanged(newVal: number) {
            if (this.patient) {
                if (newVal <= 0) {
                    this.patient.HistoryClass = 'none';
                    return;
                }
                if (newVal <= 33) {
                    this.patient.HistoryClass = 'not_strong';
                    return;
                }
                if (newVal <= 66) {
                    this.patient.HistoryClass = 'strong_family';
                    return;
                }
                if (newVal <= 100) {
                    this.patient.HistoryClass = 'strong_personal';
                }
            }
        }
    }
</script>

<style scoped>

</style>
