package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	testutil "github.com/patriciapersi/colabore-api/util"
	"github.com/stretchr/testify/assert"
)

func TestPostDispositivos(t *testing.T) {
	if err := testutil.LoadEnv(); err != nil {
		t.Fatalf("Erro ao carregar o arquivo .env: %v", err)
		t.Fatalf("%v", err)
	}
	// Define os casos de teste em uma tabela
	testCases := []struct {
		description string
		setupBody   bool
		header      map[string]string
		expected    int
	}{
		{
			description: "Inserir Dispositivo com Sucesso",
			setupBody:   true,
			header:      config.SetupHeaders(),
			expected:    http.StatusOK,
		},
		{
			description: "Tentar inserir dispositivo sem Body",
			setupBody:   false,
			header:      config.SetupHeaders(),
			expected:    http.StatusBadRequest,
		},
		{
			description: "Tentar inserir dispositivo sem header",
			setupBody:   true,
			header:      map[string]string{},
			expected:    http.StatusUnauthorized,
		},
	}

	// Itera sobre os casos de teste
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			client := config.SetupClient()
			req := client.R().
				SetHeaders(tc.header)

			// Configura o corpo da requisição se necessário
			if tc.setupBody {
				req.SetBody(config.PostDispositivosRequestBody())

			}

			resp, err := req.Post(config.BaseURL + "/agente/Dispositivos/Status")

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")

			fmt.Printf("Corpo da resposta: %s\n", resp.Body())
		})
	}
}
