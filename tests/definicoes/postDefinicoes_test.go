package main

import (
	"net/http"
	"os"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	"github.com/stretchr/testify/assert"

	testutil "github.com/patriciapersi/colabore-api/util"
)

func TestPostDefinicoes(t *testing.T) {

	defer func() {
		os.Clearenv()
	}()

	// Definindo uma tabela de casos de teste
	testCases := []struct {
		description string
		envs        map[string]string // Mapa de variáveis de ambiente a serem configuradas
		setupBody   bool
		expected    int
	}{
		{
			description: "Teste envio de Definições com sucesso",
			envs: map[string]string{
				"API_KEY": "ACTIjCHfsj1QZulkDK9oZ5bfPrlparVS44cb9o7F",
			},
			setupBody: true,
			expected:  http.StatusOK,
		},
		{
			description: "Teste envio de Definições sem body",
			envs: map[string]string{
				"API_KEY": "ACTIjCHfsj1QZulkDK9oZ5bfPrlparVS44cb9o7F",
			},
			setupBody: false,
			expected:  http.StatusBadRequest,
		},
	}

	// Iterando sobre os casos de teste
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			testutil.SetupEnv(tc.envs)

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
