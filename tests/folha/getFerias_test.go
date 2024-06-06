package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	testutil "github.com/patriciapersi/colabore-api/util"
	"github.com/stretchr/testify/assert"
)

func TestGetFerias(t *testing.T) {
	if err := testutil.LoadEnv(); err != nil {
		t.Fatalf("Erro ao carregar o arquivo .env: %v", err)
		t.Fatalf("%v", err)
	}

	testCases := []struct {
		description string
		nrInsc      string
		expected    int
	}{
		{
			description: "Busca de Ferias com Sucesso",
			nrInsc:      "63542443",
			expected:    http.StatusOK,
		},
		{
			description: "Tentativa de Buscar Férias com nrInsc vazio",
			nrInsc:      "",
			expected:    http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			client := config.SetupClient()
			url := config.BaseURL + "/agente/Ferias"

			queryParams := map[string]string{
				"nrInscEmpregador": tc.nrInsc,
			}

			resp, err := client.R().
				SetHeaders(config.SetupHeaders()).
				SetQueryParams(queryParams).
				Get(url)

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
		})
	}
}
