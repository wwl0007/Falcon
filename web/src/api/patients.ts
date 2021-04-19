import api from ".";

export interface RelativeHistoryItem {
    ID: boolean,
    CreatedAt: string,
    UpdatedAt: string,
    DeletedAt: string | null,
    Relation: string,
    Cancer: string,
    age: number,
    PatientDataID: number
}

export interface Patient {
    ID: number,
    CreatedAt: string,
    UpdatedAt: string,
    DeletedAt: string | null,
    IsPathogenic: boolean,
    Gene: string,
    HistoryClass: string,
    ethnicity: string,
    consent_approval: boolean,
    cancer_dx: boolean,
    cancer_dx_type: string,
    cancer_dx_age: number,
    known_brca: boolean,
    known_cancer: boolean,
    _method: string,
    RelativeHistory: RelativeHistoryItem[]
}

export async function fetchPatients() {
    const response = await api.get('/patients');
    return response.data;
}
