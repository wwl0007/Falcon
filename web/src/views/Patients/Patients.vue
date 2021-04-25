<template>
    <div>
        <div class="d-flex justify-content-end">
            <b-button variant="primary" @click="$router.push('/patients/edit')">Add Patient</b-button>
        </div>
        <div
            aria-controls="patients-table"
            class="d-inline-flex justify-content-between w-100"
        >
            <div>
                <div class="d-inline-flex mt-2">
                    <span class="align-self-center me-2">Page:</span>
                    <input type="number" size="3" v-model="currentPage" />
                </div>
            </div>
            <b-pagination
                :per-page="perPage"
                :total-rows="items.length"
                class="mt-2"
                :limit="10"
                :value="actualPage"
                @input="(val) => currentPage = val"
            />
        </div>
        <b-table
            striped
            hover
            :items="items"
            :fields="fields"
            :current-page="actualPage"
            :per-page="perPage"
            id="patients-table"
            :sort-by.sync="sortBy"
            :sort-desc.sync="sortDesc"
            label-sort-asc=""
            label-sort-clear=""
            label-sort-desc=""
        >
            <template #cell(Action)="row">
                <b-button @click="$router.push(`/patients/view/${row.item.ID}`)">View</b-button>
            </template>
        </b-table>
    </div>
</template>

<script lang="ts">
    import { Component, Vue, Watch } from 'vue-property-decorator'
    import { fetchPatients, Patient } from '@/api/patients.ts'
    import { Route } from "vue-router";

    @Component
    export default class Patients extends Vue {
        sortBy = 'ID';
        sortDesc = false;
        items: Patient[] = [];
        fields = [
            {
                key: 'ID',
                sortable: true
            },
            {
                key: 'Pathogenic',
                sortable: true
            },
            {
                key: 'Gene',
                sortable: true
            },
            {
                key: 'HistoryClass',
                sortable: true
            },
            {
                key: 'Ethnicity',
                sortable: true
            },
            {
                key: 'ConsentApproval',
                sortable: true
            },
            {
                key: 'CancerDX',
                sortable: true
            },
            {
                key: 'CancerDXType',
                sortable: true
            },
            {
                key: 'CancerDXAge',
                sortable: true
            },
            {
                key: 'KnownBRCA',
                sortable: true
            },
            {
                key: 'KnownCancer',
                sortable: true
            },
            'Action'
        ]

        currentPage = 1;

        get perPage() {
            return 25;
        }

        get actualPage() {
            return this.currentPage || 1;
        }

        async mounted() {
            this.items = await fetchPatients();

            this.updatePageFromURL(this.$route);
        }

        @Watch('currentPage')
        updatePageInUrl() {
            // eslint-disable-next-line @typescript-eslint/no-empty-function
            this.$router.replace(`?page=${this.actualPage}`).catch(() => {});
        }

        @Watch('$route')
        updatePageFromURL(newRoute: Route) {
            const page = newRoute.query.page as string | undefined;
            if (page !== undefined) {
                this.currentPage = parseInt(page);
            }
        }
    }
</script>

<style scoped>

</style>
