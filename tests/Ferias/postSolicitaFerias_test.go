//Aguardando script que faça inclusão da solicitação de férias pelo próprio teste

package main

import (
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	testutil "github.com/patriciapersi/colabore-api/util"
)

func TestPostSolicitaFerias(t *testing.T) {
	if err := testutil.LoadEnv(); err != nil {
		t.Fatalf("Erro ao carregar o arquivo .env: %v", err)
		t.Fatalf("%v", err)
	}
	testCases := []struct {
		description string
		setupBody   bool
		header      map[string]string
		// expected     int
		// expectedDesc string
	}{
		{
			description: "Solicitar férias com sucesso",
			setupBody:   true,
			header:      config.SetupHeaders(),
			// expected:     http.StatusOK,
			// expectedDesc: "Sucesso",
		},
	}
	// Itera sobre os casos de teste
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			api := config.SetupApi()
			req := api.Client.R().SetHeaders(tc.header)

			// Configura o corpo da requisição se necessário
			if tc.setupBody {
				req.SetBody(config.PostSolicitaFeriasRequestBody())
			}

			//resp, err := req.Post(api.Endpoints["POSTferias"])

			// assert.NoError(t, err, "Erro ao fazer a requisição")
			// assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			// assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}

}
