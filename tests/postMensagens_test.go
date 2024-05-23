package main

import (
	"net/http"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/tonnarruda/API_Colabore/config"
)

func TestPostMensagens(t *testing.T) {
	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load("c:\\workspace\\colabore-api\\.env"); err != nil {
		t.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Definindo uma tabela de casos de teste
	testCases := []struct {
		description string
		setupBody   bool
		expected    int
	}{
		{
			description: "Teste envio de mensagem",
			setupBody:   true,
			expected:    http.StatusOK,
		},
		{
			description: "Teste envio de mensagem",
			setupBody:   false,
			expected:    http.StatusBadRequest,
		},
	}

	// Iterando sobre os casos de teste
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			client := config.SetupClient()
			req := client.R().
				SetHeaders(config.SetupHeaders())

			// Configura o corpo da requisição se necessário
			if tc.setupBody {
				req.SetBody(config.MensagensRequestBody())

			}

			resp, err := req.Post(config.BaseURL + "/agente/mensagem")

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
		})
	}
}
