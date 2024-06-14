package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	"github.com/stretchr/testify/assert"
)

func TestDispositivos(t *testing.T) {

	testCases := []struct {
		description  string
		header       map[string]string
		status       string
		expected     int
		expectedDesc string
	}{
		{
			description:  "Busca de dispositivos com Sucesso",
			header:       config.SetupHeadersAgente(),
			status:       "0",
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Busca de dispositivos com Status Inexistente",
			header:       config.SetupHeadersAgente(),
			status:       "-1",
			expected:     http.StatusBadRequest,
			expectedDesc: "Status' possui um intervalo de valores que não inclui '-1'.",
		},
		{
			description:  "Busca de dispositivos com Status Vazio",
			header:       map[string]string{},
			status:       "",
			expected:     http.StatusUnauthorized,
			expectedDesc: "Unauthorized",
		},
	}

	// Itera sobre os casos de teste
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			api := config.SetupApi()
			queryParams := map[string]string{
				"status": tc.status,
			}

			resp, err := api.Client.R().
				SetHeaders(tc.header).
				SetQueryParams(queryParams).
				Get(api.EndpointsAgente["Dispositivos"])

			// Verifica os resultados do teste
			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}
}
