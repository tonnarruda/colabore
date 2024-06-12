package structs

// GeridoLista representa um funcionário gerido por um gestor na ListaGeridos
type ListaGeridos struct {
	CPF              string `json:"CPF"`
	Matricula        string `json:"Matricula"`
	NrInscEmpregador string `json:"NrInscEmpregador"`
	NomeFantasia     string `json:"NomeFantasia"`
}

// GeridoGeridos representa um funcionário gerido por um gestor na Geridos
type Geridos struct {
	CPF       string `json:"CPF"`
	Matricula string `json:"Matricula"`
}

// Gestor representa um gestor com sua lista de geridos
type Gestor struct {
	NrInscEmpregador string         `json:"NrInscEmpregador"`
	CPFGestor        string         `json:"CPFGestor"`
	MatriculaGestor  string         `json:"MatriculaGestor"`
	ListaGeridos     []ListaGeridos `json:"ListaGeridos"`
	Geridos          []Geridos      `json:"Geridos"`
}
