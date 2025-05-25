package response

type Metric struct {
	ID             uint   `json:"id"`
	Data           string `json:"data"`
	CorretorID     uint   `json:"id_corretor"`
	LeadsRecebidos int    `json:"leads_recebidos"`
	Ligacoes       int    `json:"ligacoes"`
	Espontaneo     int    `json:"espontaneo"`
	Captacoes      int    `json:"captacoes"`
	Visitas        int    `json:"visitas"`
	Negociacoes    int    `json:"negociacoes"`
	Propostas      int    `json:"propostas"`
	Vendas         int    `json:"vendas"`
	Observacao     string `json:"observacao"`
}
