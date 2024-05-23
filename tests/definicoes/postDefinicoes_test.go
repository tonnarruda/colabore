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
	// Carrega as variáveis de ambiente do arquivo .env
	if err := testutil.LoadEnv(); err != nil {
		t.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

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
				"API_KEY": os.Getenv("API_KEY"),
			},
			setupBody: true,
			expected:  http.StatusOK,
		},
		{
			description: "Teste envio de Definições sem body",
			envs: map[string]string{
				"API_KEY": os.Getenv("API_KEY"),
			},
			setupBody: false,
			expected:  http.StatusBadRequest,
		},
		{
			description: "Teste com API_KEY vazia",
			envs: map[string]string{
				"API_KEY": "",
			},
			setupBody: true,
			expected:  http.StatusForbidden,
		},
		{
			description: "Teste com CNPJ_LICENCIADO, COLABORE_SIGNATURE_EXPIRATION e COLABORE_SIGNATURE vazios",
			envs: map[string]string{
				"CNPJ_LICENCIADO":               "",
				"COLABORE_SIGNATURE_EXPIRATION": "",
				"COLABORE_SIGNATURE":            "",
			},
			setupBody: true,
			expected:  http.StatusUnauthorized,
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
