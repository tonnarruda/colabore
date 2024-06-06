package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	testutil "github.com/patriciapersi/colabore-api/util"
	"github.com/stretchr/testify/assert"
)

func TestGetColaboradorPremium(t *testing.T) {
	if err := testutil.LoadEnv(); err != nil {
		t.Fatalf("Erro ao carregar o arquivo .env: %v", err)
		t.Fatalf("%v", err)
	}

	testCases := []struct {
		description      string
		cpf              string
		nrInscEmpregador string
		expected         int
		expectedDesc     string
	}{
		{
			description:      "Busca de candidato - Não Encontrado",
			cpf:              "60515860409",
			nrInscEmpregador: "10821992",
			expected:         http.StatusBadRequest,
			expectedDesc:     "Candidato com dados preenchidos não encontrado",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			client := config.SetupClient()
			url := config.BaseURL + "/agente/Candidato"

			queryParams := map[string]string{
				"CPF":              tc.cpf,
				"NrInscEmpregador": tc.nrInscEmpregador,
			}

			resp, err := client.R().
				SetHeaders(config.SetupHeaders()).
				SetQueryParams(queryParams).
				Get(url)

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}
}
