//Aguardando script que faça inclusão da solicitação de férias pelo próprio teste

package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	testutil "github.com/patriciapersi/colabore-api/util"
	"github.com/stretchr/testify/assert"
)

// Esta função é a precondição de ter a solicitação de ferias enviada pelo app
func precondition() {
	api := config.SetupApi()
	api.Client.R().
		SetBody(config.PostSolicitaFeriasAPPRequestBody()).
		SetHeaders(config.SetupHeadersApp()).
		Post(api.EndpointsApp["POSTappferias"])
}

func TestPostSolicitaFerias_(t *testing.T) {
	testutil.LoadEnv()

	testCases := []struct {
		description  string
		setupBody    bool
		header       map[string]string
		expected     int
		expectedDesc string
		precondition func()
	}{
		{
			description:  "Test Success",
			setupBody:    true,
			header:       config.SetupHeadersAgente(),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
			precondition: precondition,
		},
		{
			description:  "Test BadRequest",
			setupBody:    false,
			header:       config.SetupHeadersAgente(),
			expected:     http.StatusBadRequest,
			expectedDesc: "Corpo da requisição não contém solicitações de férias",
			precondition: precondition,
		},
		{
			description:  "Test Unauthorized",
			setupBody:    true,
			header:       map[string]string{},
			expected:     http.StatusUnauthorized,
			expectedDesc: "Unauthorized",
			precondition: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// Executa a precondição se ela não for nil
			if tc.precondition != nil {
				tc.precondition()
			}

			api := config.SetupApi()

			var body interface{}
			if tc.setupBody {
				body = config.PostSolicitaFeriasRequestBody()

			}

			resp, err := api.Client.R().
				SetHeaders(tc.header).
				SetBody(body).
				Post(api.EndpointsAgente["Ferias"])

			assert.NoError(t, err, "Erro ao fazer a requisição para %s", tc.description)
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado para %s", tc.description)
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}
}
