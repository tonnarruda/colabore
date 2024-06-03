package main

import (
	"net/http"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/patriciapersi/colabore-api/config"
	"github.com/stretchr/testify/assert"
)

func TestGetImagem(t *testing.T) {

	testCases := []struct {
		description string
		nome        string
		expected    int
	}{
		{
			description: "Buscar Imagem com Nome Inexistente",
			nome:        faker.Word(),
			expected:    http.StatusBadRequest,
		},
		{
			description: "Buscar Imagem com Nome Vazio",
			nome:        "",
			expected:    http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			client := config.SetupClient()
			url := config.BaseURL + "/agente/Imagem"

			queryParams := map[string]string{
				"Nome": tc.nome,
			}

			resp, err := client.R().
				SetHeaders(config.SetupHeaders()).
				SetQueryParams(queryParams).
				Get(url)

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
		})
	}
}
