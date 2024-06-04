package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	"github.com/stretchr/testify/assert"
)

func TestPostDefinicoes(t *testing.T) {
	testCases := []struct {
		description string
		setupBody   bool
		expected    int
	}{
		{
			description: "Teste envio de Definições com sucesso",
			setupBody:   true,
			expected:    http.StatusOK,
		},
		{
			description: "Teste envio de Definições sem body",
			setupBody:   false,
			expected:    http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			client := config.SetupClient()
			req := client.R().
				SetHeaders(config.SetupHeaders())

			// Configura o corpo da requisição se necessário
			if tc.setupBody {
				req.SetBody(config.DefinicoesRequestBody())
			}

			resp, err := req.Post(config.BaseURL + "/agente/licenciado/definicoes")

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
		})
	}
}
