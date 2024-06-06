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
		description string
		nrInsc      string
		lastupdate  string
		expected    int
	}{
		{
			description: "Tentativa de buscar colaborador Freemium - Não Encontrado",
			nrInsc:      "63542443",
			lastupdate:  time.Now().Format("2006-01-02 15:04:05.000 -0700"),
			expected:    http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			client := config.SetupClient()
			url := config.BaseURL + "/agente/Colaborador/Fremium"

			queryParams := map[string]string{
				"nrInscEmpregador": tc.nrInsc,
				"lastupdate":       tc.lastupdate,
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
