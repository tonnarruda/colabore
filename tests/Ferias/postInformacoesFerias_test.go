package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	testutil "github.com/patriciapersi/colabore-api/util"
	"github.com/stretchr/testify/assert"
)

//Inserimos as definições para que o usuário possa solicitar férias

func TestPostDefinicoesFerias(t *testing.T) {
	if err := testutil.LoadEnv(); err != nil {
		t.Fatalf("Erro ao carregar o arquivo .env: %v", err)
		t.Fatalf("%v", err)
	}
	testCases := []struct {
		description string
		setupBody   bool
		header      map[string]string
		expected    int
	}{
		{
			description: "Teste envio de Definições com sucesso",
			setupBody:   true,
			header:      config.SetupHeaders(),
			expected:    http.StatusOK,
		},
		{
			description: "Teste envio de Definições sem body",
			setupBody:   false,
			header:      config.SetupHeaders(),
			expected:    http.StatusBadRequest,
		},
		{
			description: "Teste envio de Definições com header vazio",
			setupBody:   true,
			header:      map[string]string{},
			expected:    http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			client := config.SetupClient()
			req := client.R().
				SetHeaders(tc.header)

			// Configura o corpo da requisição se necessário
			if tc.setupBody {
				req.SetBody(config.PostInformacoesFeriasEmpregadoRequestBody())
			}

			resp, err := req.Post(config.BaseURL + "/agente/ferias/informacoes")
			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")

		})
	}
}
