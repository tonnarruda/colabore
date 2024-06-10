package main

import (
	"net/http"
	"testing"
	"time"

	"github.com/patriciapersi/colabore-api/config"
	testutil "github.com/patriciapersi/colabore-api/util"
	"github.com/stretchr/testify/assert"
)

func TestGetColaboradorFremium(t *testing.T) {
	if err := testutil.LoadEnv(); err != nil {
		t.Fatalf("Erro ao carregar o arquivo .env: %v", err)
		t.Fatalf("%v", err)
	}

	testCases := []struct {
		description  string
		nrInsc       string
		header       map[string]string
		lastupdate   string
		expected     int
		expectedDesc string
	}{
		{
			description:  "Tentativa de buscar colaborador Freemium - Não Encontrado",
			nrInsc:       "10821992",
			header:       config.SetupHeadersAgente(),
			lastupdate:   time.Now().Format("2006-01-02 15:04:05.000 -0700"),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			api := config.SetupApi()

			queryParams := map[string]string{
				"nrInscEmpregador": tc.nrInsc,
				"lastupdate":       tc.lastupdate,
			}

			resp, err := api.Client.R().
				SetHeaders(tc.header).
				SetQueryParams(queryParams).
				Get(api.EndpointsAgente["GETcolaboradorFreemium"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}
}
