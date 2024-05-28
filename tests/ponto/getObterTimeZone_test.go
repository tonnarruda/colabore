package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	"github.com/stretchr/testify/assert"

	testutil "github.com/patriciapersi/colabore-api/util"
)

func TestGetTimeZone(t *testing.T) {
	// Carrega as variáveis de ambiente do arquivo .env
	if err := testutil.LoadEnv(); err != nil {
		t.Fatalf("%v", err)
	}

	// Define os casos de teste em uma tabela
	testCases := []struct {
		description string
		latitude    string
		longitude   string
		expected    int
	}{
		{
			description: "Envio de latitude e longitude válidas",
			latitude:    "-22.9519",
			longitude:   "-43.2105",
			expected:    http.StatusOK,
		},
		{
			description: "Envio de latitude e longitude inválidas",
			latitude:    "latitude",
			longitude:   "longitude",
			expected:    http.StatusBadRequest,
		},
		{
			description: "Envio de latitude e longitude vazias",
			latitude:    "",
			longitude:   "",
			expected:    http.StatusBadRequest,
		},
	}

	// Itera sobre os casos de teste
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// Configura o cliente e a URL
			client := config.SetupClient()
			url := config.BaseURL + "/agente/Ponto/ObterTimeZone"
			url = fmt.Sprintf("%s?latitude=%s&longitude=%s", url, tc.latitude, tc.longitude)

			// Faz a requisição GET
			resp, err := client.R().
				SetHeaders(config.SetupHeaders()).
				Get(url)

			// Verifica os resultados do teste
			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")

			fmt.Println("Resposta da requisição:", resp.String())
		})
	}
}
