//Aguardando script que faça inclusão da solicitação de férias pelo próprio teste

package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	testutil "github.com/patriciapersi/colabore-api/util"
	"github.com/stretchr/testify/assert"
)

func TestDeleteAssinatura(t *testing.T) {
	testutil.LoadEnv()

	testCases := []struct {
		description  string
		setupBody    bool
		header       map[string]string
		expected     int
		expectedDesc string
	}{
		{
			description:  "Deletar Assinatura com Sucesso",
			setupBody:    true,
			header:       config.SetupHeadersAgente(),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Tentativa de Deletar Assinatura sem o body no corpo da requisição",
			setupBody:    false,
			header:       config.SetupHeadersAgente(),
			expected:     http.StatusBadRequest,
			expectedDesc: "Corpo da requisição não contém assinatura",
		},
		{
			description:  "Tentativa de Liberar Assinatura sem o header de Autenticação",
			setupBody:    true,
			header:       map[string]string{},
			expected:     http.StatusUnauthorized,
			expectedDesc: "Unauthorized",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			api := config.SetupApi()
			var body interface{}
			if tc.setupBody {
				body = config.DeleteAssinaturaRequestBody()
			}

			resp, err := api.Client.R().
				SetHeaders(tc.header).
				SetBody(body).
				Delete(api.EndpointsAgente["Assinatura"])

			assert.NoError(t, err, "Erro ao fazer a requisição para %s", tc.description)
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado para %s", tc.description)
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}
}
