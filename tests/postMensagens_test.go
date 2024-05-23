package main

import (
	"net/http"
	"os"
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

	// Captura as chaves das variáveis de ambiente a partir dos headers configurados
	envKeys := config.GetEnvKeysFromHeaders()

	// Captura o estado original das variáveis de ambiente
	originalEnv := config.CaptureOriginalEnv(envKeys)

	// Definindo uma tabela de casos de teste
	testCases := []struct {
		description string
		envs        map[string]string // Mapa de variáveis de ambiente a serem configuradas
		setupBody   bool
		expected    int
	}{
		{
			description: "Teste envio de mensagem",
			envs: map[string]string{
				"API_KEY":                       os.Getenv("API_KEY"),
				"CNPJ_LICENCIADO":               os.Getenv("CNPJ_LICENCIADO"),
				"COLABORE_SIGNATURE_EXPIRATION": os.Getenv("COLABORE_SIGNATURE_EXPIRATION"),
				"COLABORE_SIGNATURE":            os.Getenv("COLABORE_SIGNATURE"),
			},
			setupBody: true,
			expected:  http.StatusOK,
		},
	}

	// Iterando sobre os casos de teste
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			// Restaurar variáveis de ambiente após cada teste
			t.Cleanup(func() { config.RestoreEnv(originalEnv) })

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
