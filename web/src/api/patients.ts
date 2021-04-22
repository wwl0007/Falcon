import api from ".";

export interface RelativeHistoryItem {
    ID: boolean,
    CreatedAt: string,
    UpdatedAt: string,
    DeletedAt: string | null,
    Relation: string,
    Cancer: string,
    Age: number,
    PatientDataID: number
}

export interface Patient {
    ID: number,
    CreatedAt: string,
    UpdatedAt: string,
    DeletedAt: string | null,
    Pathogenic: string,
    Gene: string,
    HistoryClass: string,
    Ethnicity: string,
    ConsentApproval: string,
    CancerDX: string,
    CancerDXType: string,
    CancerDXAge: number,
    KnownBRCA: string,
    KnownCancer: string,
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
