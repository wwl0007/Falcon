import api from ".";

export interface RelativeHistoryItem {
    ID: number,
    CreatedAt: string,
    UpdatedAt: string,
    DeletedAt: string | null,
    Relation: string,
    Cancer: string,
    Age: number,
    PatientDataID: number
}

export interface BasicPatientInfo {
    Pathogenic: string,
    Gene: string,
    Ethnicity: string,
    ConsentApproval: string,
    CancerDX: string,
    CancerDXType: string,
    CancerDXAge: number,
    KnownBRCA: string,
    KnownCancer: string
}

export interface Patient extends BasicPatientInfo {
    ID: number,
    CreatedAt: string,
    UpdatedAt: string,
    DeletedAt: string | null,
    HistoryClass: string,
    RelativeHistory: RelativeHistoryItem[]
}

export async function fetchPatients() {
    const response = await api.get('/patients');
    return response.data;
}

export async function fetchPatientById(patientId: number) {
    const response = await api.get(`/patients/${patientId}`);
    return response.data;
}

export async function deleteRelativeHistory(id: number) {
    const response = await api.delete(`/relativeHistory/${id}`);
    return response.data;
}

export async function updateRelativeHistory(value: RelativeHistoryItem) {
    const response = await api.put(
        `/relativeHistory`,
        {
            ID: value.ID,
            Relation: value.Relation,
            Cancer: value.Cancer,
            Age: value.Age,
            PatientDataID: value.PatientDataID
        }
    );

    return response.data;
}

export async function createPatient(patient: BasicPatientInfo) {
    const response = await api.put(
        `/patients`,
        {
            Pathogenic: patient.Pathogenic,
            Gene: patient.Gene,
            full_history: {
                ethnicity: patient.Ethnicity,
                consent_approval: patient.ConsentApproval,
                cancer_dx: patient.CancerDX,
                cancer_dx_type: patient.CancerDXType,
                cancer_dx_age: patient.CancerDXAge,
                known_brca: patient.KnownBRCA,
                known_cancer: patient.KnownCancer
            }
        }
    );

    return response.data;
}
