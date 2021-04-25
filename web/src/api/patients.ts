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
    Pathogenic: boolean,
    Gene: string,
    Ethnicity: string,
    ConsentApproval: boolean,
    CancerDX: boolean,
    CancerDXType: string,
    CancerDXAge: number,
    KnownBRCA: boolean,
    KnownCancer: boolean,
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
            patient
        }
    );

    return response.data;
}
