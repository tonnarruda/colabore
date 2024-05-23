package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	"github.com/stretchr/testify/assert"

	testutil "github.com/patriciapersi/colabore-api/util"
)

func TestGetDispositivos(t *testing.T) {
	// Carrega as variáveis de ambiente do arquivo .env
	if err := testutil.LoadEnv(); err != nil {
		t.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Define os casos de teste em uma tabela
	testCases := []struct {
		description string
		status      interface{}
		expected    int
	}{
		{
			description: "Busca de dispositivos com Sucesso",
			status:      0,
			expected:    http.StatusOK,
		},
		{
			description: "Busca de dispositivos com Status Inexistente",
			status:      -1,
			expected:    http.StatusBadRequest,
		},
		{
			description: "Busca de dispositivos com Status Nulo",
			status:      nil,
			expected:    http.StatusBadRequest,
		},
		{
			description: "Busca de dispositivos com Status Vazio",
			status:      "",
			expected:    http.StatusBadRequest,
		},
	}

	// Itera sobre os casos de teste
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// Configura o cliente e a URL
			client := config.SetupClient()
			url := config.BaseURL + "/agente/Dispositivos"
			url = fmt.Sprintf("%s?status=%d", url, tc.status)

			// Faz a requisição GET
			resp, err := client.R().
				SetHeaders(config.SetupHeaders()).
				Get(url)

			// Verifica os resultados do teste
			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
		})
	}
}
