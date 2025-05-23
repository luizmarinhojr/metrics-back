export interface MetricRequest {
    data: string;
    nome: string;
    leadsRecebidos: number;
    ligacoes: number;
    espontaneo: number;
    captacoes: number;
    visitas: number;
    negociacoes: number;
    propostas: number;
    vendas: number;
    observacao: string;
}

export interface MetricResponse {
    id: number;
    data: string;
    nome: string;
    leadsRecebidos: number;
    ligacoes: number;
    espontaneo: number;
    captacoes: number;
    visitas: number;
    negociacoes: number;
    propostas: number;
    vendas: number;
    observacao: string;
}
