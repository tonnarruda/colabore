package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	testutil "github.com/patriciapersi/colabore-api/util"
	"github.com/stretchr/testify/assert"
)

func TestFerias(t *testing.T) {
	testutil.LoadEnv()

	testCases := []struct {
		description  string
		nrInsc       string
		header       map[string]string
		expected     int
		expectedDesc string
	}{
		{
			description:  "Busca de Ferias com Sucesso",
			nrInsc:       "63542443",
			header:       config.SetupHeadersAgente(),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Tentativa de Buscar Férias com nrInsc vazio",
			nrInsc:       "",
			header:       config.SetupHeadersAgente(),
			expected:     http.StatusBadRequest,
			expectedDesc: "'Nr Insc Empregador' deve ser informado.",
		},
		{
			description:  "Tentativa de Buscar Férias com nrInsc vazio",
			nrInsc:       "63542443",
			header:       map[string]string{},
			expected:     http.StatusUnauthorized,
			expectedDesc: "Unauthorized",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			api := config.SetupApi()
			queryParams := map[string]string{
				"nrInscEmpregador": tc.nrInsc,
			}

			resp, err := api.Client.R().
				SetHeaders(tc.header).
				SetQueryParams(queryParams).
				Get(api.EndpointsAgente["Ferias"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}
}
