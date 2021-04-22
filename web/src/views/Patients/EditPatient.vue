<template>
    <div class="border rounded bg-body d-flex justify-content-center p-4">
        <div class="d-flex flex-column">
            <div>
                <b-button @click="$router.back()">Back</b-button>
            </div>

            <!-- form inputs -->
            <div class="p-4 mt-2">
                <h3 class="mt-2 text-center">Add Patient</h3>

                <div class="form-input">
                    Is Pathogenic?
                    <b-checkbox v-model="isPathogenic" />
                </div>
                <div class="form-input">
                    Gene
                    <b-checkbox-group
                        v-model="selectedLetters"
                        :options="lettersOptions"
                    />
                </div>
                <div class="form-input">
                    Ethnicity
                    <b-input v-model="ethnicity" />
                </div>
                <div class="form-input">
                    Consent Approval
                    <b-checkbox v-model="consentApproval" />
                </div>
                <div class="form-input">
                    Cancer DX
                    <b-checkbox v-model="cancerDX" />
                </div>
                <div class="form-input">
                    Cancer DX Type
                    <b-input v-model="cancerDXType" />
                </div>
                <div class="form-input">
                    Cancer DX Age
                    <b-input v-model="cancerDXAge" type="number" />
                </div>
                <div class="form-input">
                    Known BRCA
                    <b-checkbox v-model="knownBRCA" />
                </div>
                <div class="form-input">
                    Known Cancer
                    <b-checkbox v-model="knownCancer" />
                </div>

                <FamilyHistoryViewer v-model="relativeHistory" />

                <div class="d-flex justify-content-center">
                    <b-button variant="primary">Submit</b-button>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
    import { Component, Vue } from 'vue-property-decorator'
    import { RelativeHistoryItem } from '@/api/patients'
    import FamilyHistoryViewer from '@/components/FamilyHistory/FamilyHistoryViewer.vue'

    @Component({
        components: { FamilyHistoryViewer }
    })
    export default class EditPatient extends Vue {
        isPathogenic = false;
        ethnicity = "";
        consentApproval = false;
        cancerDX = false;
        cancerDXType = "";
        cancerDXAge: number | null = null;
        knownBRCA = false;
        knownCancer = false;
        relativeHistory: RelativeHistoryItem[] = [];

        lettersOptions = ["A", "C", "G", "T"];
        selectedLetters = [];

        get gene() {
            return this.selectedLetters.join("");
        }

        mounted() {
            console.log(this.$router.currentRoute.params);
        }
    }
</script>

<style scoped>
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
</style>
