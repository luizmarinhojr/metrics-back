import { BrokerResponse } from "./broker";

export interface MetricRequest {
    data: string;
    leads_recebidos?: number;
    ligacoes?: number;
    espontaneo?: number;
    captacoes?: number;
    visitas?: number;
    negociacoes?: number;
    propostas?: number;
    vendas?: number;
    agendamentos?: number;
}

export interface MetricResponse {
    id: number;
    data: string;
    corretor: BrokerResponse;
    leads_recebidos: number;
    ligacoes: number;
    espontaneo: number;
    captacoes: number;
    visitas: number;
    negociacoes: number;
    propostas: number;
    vendas: number;
    agendamentos: number;
    created_at: string;
    updated_at: string;
}