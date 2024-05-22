package main

import (
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/tonnarruda/API_Colabore/config"

	testutil "github.com/tonnarruda/API_Colabore/util"
)

func TestPostDefinicoes(t *testing.T) {
	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load("c:\\workspace\\colabore\\.env"); err != nil {
		t.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

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

// package main

// import (
// 	"os"
// 	"testing"

// 	"github.com/joho/godotenv"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/tonnarruda/API_Colabore/config"
// )

// func TestPostColabore(t *testing.T) {
// 	if err := godotenv.Load("c:\\workspace\\colabore\\.env"); err != nil {
// 		t.Fatalf("Erro ao carregar o arquivo .env: %v", err)
// 	}

// 	client := config.SetupClient()
// 	body := config.DefinicoesRequestBody()
// 	url := config.BaseURL + "/agente/licenciado/definicoes"
// 	resp, err := client.R().
// 		SetHeaders(config.SetupHeaders()).
// 		SetBody(body).
// 		Post(url)

// 	assert.NoError(t, err, "Erro ao fazer a requisição")
// 	assert.Equal(t, 200, resp.StatusCode(), "Status de resposta inesperado")
// }

// func TestPostColabore401(t *testing.T) {
// 	if err := godotenv.Load("c:\\workspace\\colabore\\.env"); err != nil {
// 		t.Fatalf("Erro ao carregar o arquivo .env: %v", err)
// 	}
// 	os.Setenv("API_KEY", "")

// 	client := config.SetupClient()
// 	body := config.DefinicoesRequestBody()
// 	url := config.BaseURL + "/agente/licenciado/definicoes"
// 	resp, err := client.R().
// 		SetHeaders(config.SetupHeaders()).
// 		SetBody(body).
// 		Post(url)

// 	assert.NoError(t, err, "Erro ao fazer a requisição")
// 	assert.Equal(t, 403, resp.StatusCode(), "Status de resposta inesperado")
// }
