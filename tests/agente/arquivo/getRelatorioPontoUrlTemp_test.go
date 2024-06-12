package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	testutil "github.com/patriciapersi/colabore-api/util"
	"github.com/stretchr/testify/assert"
)

func TestGetRelatorioPontoUrlTemp(t *testing.T) {
	testutil.LoadEnv()

	testCases := []struct {
		description      string
		cpf              string
		nrInscEmpregador string
		header           map[string]string
		matricula        string
		expected         int
		expectedDesc     string
	}{
		{
			description:      "Buscar Relatorio de Ponto com Sucesso",
			cpf:              "60515860409",
			nrInscEmpregador: "10821992",
			header:           config.SetupHeadersAgente(),
			matricula:        "000031",
			expected:         http.StatusOK,
			expectedDesc:     "Sucesso",
		},
		{
			description:      "Buscar Relatorio de Ponto nrInsc Vazio",
			cpf:              "60515860409",
			nrInscEmpregador: "",
			header:           config.SetupHeadersAgente(),
			matricula:        "000031",
			expected:         http.StatusBadRequest,
			expectedDesc:     "Caminho,NrInscEmpregador",
		},
		{
			description:      "Buscar Relatorio de Ponto CPF Vazio",
			cpf:              "",
			nrInscEmpregador: "10821992",
			header:           config.SetupHeadersAgente(),
			matricula:        "000031",
			expected:         http.StatusBadRequest,
			expectedDesc:     "Caminho,CPF",
		},
		{
			description:      "Buscar Relatorio de Ponto Matricula Vazio",
			cpf:              "60515860409",
			nrInscEmpregador: "10821992",
			header:           config.SetupHeadersAgente(),
			matricula:        "",
			expected:         http.StatusBadRequest,
			expectedDesc:     "Caminho,Matricula",
		},
		{
			description:      "Buscar Relatorio de Ponto com Sucesso",
			cpf:              "60515860409",
			nrInscEmpregador: "10821992",
			header:           map[string]string{},
			matricula:        "000031",
			expected:         http.StatusUnauthorized,
			expectedDesc:     "Unauthorized",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			api := config.SetupApi()
			queryParams := map[string]string{
				"CPF":              tc.cpf,
				"NrInscEmpregador": tc.nrInscEmpregador,
				"Matricula":        tc.matricula,
				"AnoMes":           "202401",
			}

			resp, err := api.Client.R().
				SetHeaders(tc.header).
				SetQueryParams(queryParams).
				Get(api.EndpointsAgente["ArquivoRelatorioPontoURLTemporaria"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}
}
