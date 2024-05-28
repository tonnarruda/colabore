package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	"github.com/stretchr/testify/assert"
)

func TestGetColaboradorPremium(t *testing.T) {

	testCases := []struct {
		description string
		expected    int
	}{
		{
			description: "Busca de Colaborador Premium com Sucesso",
			expected:    http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			client := config.SetupClient()
			url := config.BaseURL + "/agente/Colaborador/Premium"

			resp, err := client.R().
				SetHeaders(config.SetupHeaders()).
				Get(url)

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
		})
	}
}
