<template>
    <div class="border rounded bg-body p-4">
        <div class="mb-4">
            <b-button @click="$router.back()">Back</b-button>
        </div>
        <div class="d-flex justify-content-center">
            <div style="min-width: 50%">
                <!-- form inputs -->
                <div class="p-4 mt-2">
                    <h3 class="mt-2 text-center">Add Patient</h3>

                    <div class="form-input">
                        Is Pathogenic?
                        <b-input v-model="isPathogenic" />
                    </div>
                    <div class="form-input">
                        <b-input-group prepend="Gene">
                            <b-form-input v-model="gene" />
                        </b-input-group>
                    </div>
                    <div class="form-input">
                        <b-input-group prepend="Ethnicity">
                            <b-form-input v-model="ethnicity" />
                        </b-input-group>
                    </div>
                    <div class="form-input">
                        Consent Approval
                        <b-input v-model="consentApproval"/>
                    </div>
                    <div class="form-input">
                        Cancer DX
                        <b-input v-model="cancerDX" />
                    </div>
                    <div class="form-input">
                        <b-input-group prepend="Cancer DX Type">
                            <b-form-input v-model="cancerDXType" />
                        </b-input-group>
                    </div>
                    <div class="form-input">
                        Cancer DX Age
                        <b-input v-model="cancerDXAge" type="number" />
                    </div>
                    <div class="form-input">
                        Known BRCA
                        <b-input v-model="knownBRCA" />
                    </div>
                    <div class="form-input">
                        Known Cancer
                        <b-input v-model="knownCancer" />
                    </div>

                    <div class="d-flex justify-content-center">
                        <b-button variant="primary" @click="submitForm">Submit</b-button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
    import { Component, Vue } from 'vue-property-decorator'
    import { createPatient, Patient } from '@/api/patients'
    import FamilyHistoryViewer from '@/components/FamilyHistory/FamilyHistoryViewer.vue'

    @Component({
        components: { FamilyHistoryViewer }
    })
    export default class AddPatient extends Vue {
        isPathogenic = "FALSE";
        ethnicity = "";
        consentApproval = "FALSE";
        cancerDX = "FALSE";
        cancerDXType = "";
        cancerDXAge = 0;
        knownBRCA = "FALSE";
        knownCancer = "FALSE";
        gene = "";

        async submitForm() {
            const result = await createPatient(
                {
                    Pathogenic: this.isPathogenic,
                    Gene: this.gene,
                    Ethnicity: this.ethnicity,
                    ConsentApproval: this.consentApproval,
                    CancerDX: this.cancerDX,
                    CancerDXType: this.cancerDXType,
                    CancerDXAge: this.cancerDXAge,
                    KnownBRCA: this.knownBRCA,
                    KnownCancer: this.knownCancer
                }
            ) as Patient;

            await this.$router.push(`/patients/view/${result.ID}`);
        }
    }
</script>

<style>
    .new-patient-form {
        border: 2px solid #2c3e50;
        border-radius: 8px;
    }

    div {
        font-size: 18px;
    }

    .form-input {
        margin-top: 10px;
        margin-bottom: 20px;
    }

    input[type='checkbox'] {
        width: 1.5rem;
        height: 1.5rem;
    }
</style>
