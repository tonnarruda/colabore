package main

import (
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/patriciapersi/colabore-api/config"
	testutil "github.com/patriciapersi/colabore-api/util"
	"github.com/stretchr/testify/assert"
)

// Configuração comum para todos os testes
func setupTest() (string, *resty.Client) {
	return config.BaseURL + "/agente/mensagem", config.SetupClient()
}

// Testa a exclusão de uma mensagem
func TestDeleteMensagens(t *testing.T) {
	if err := testutil.LoadEnv(); err != nil {
		t.Fatalf("Erro ao carregar o arquivo .env: %v", err)
		t.Fatalf("%v", err)
	}
	tests := []struct {
		description string
		id          string
		expected    int
	}{
		{
			description: "Mensagem existente",
			id: func() string {
				url, client := setupTest()
				requestBody := config.MensagensRequestBody()
				id := requestBody["ID"].(string)
				reqPost := client.R().
					SetHeaders(config.SetupHeaders()).
					SetBody(requestBody)

				reqPost.Post(url)
				return id
			}(),
			expected: http.StatusOK,
		},
		{
			description: "ID inexistente",
			id:          uuid.New().String(),
			expected:    http.StatusOK,
		},
		{
			description: "ID inexistente",
			id:          "",
			expected:    http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			url, client := setupTest()

			deleteBody := config.DeleteMensagensRequestBody(tt.id)
			req := client.R().
				SetHeaders(config.SetupHeaders()).
				SetBody(deleteBody)

			resp, err := req.Delete(url)

			assert.NoError(t, err, "Erro ao fazer a requisição para %s", tt.description)
			assert.Equal(t, tt.expected, resp.StatusCode(), "Status de resposta inesperado para %s", tt.description)
		})
	}
}
