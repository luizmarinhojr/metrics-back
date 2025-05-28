package response

import "time"

type Metric struct {
	ID             uint       `json:"id"`
	Data           string     `json:"data"`
	Corretor       Broker     `json:"corretor"`
	LeadsRecebidos int        `json:"leads_recebidos"`
	Ligacoes       int        `json:"ligacoes"`
	Espontaneo     int        `json:"espontaneo"`
	Captacoes      int        `json:"captacoes"`
	Visitas        int        `json:"visitas"`
	Negociacoes    int        `json:"negociacoes"`
	Propostas      int        `json:"propostas"`
	Vendas         int        `json:"vendas"`
	Observacao     string     `json:"observacao"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}
