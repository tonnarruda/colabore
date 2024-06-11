//Aguardando script que faça inclusão da solicitação de férias pelo próprio teste

package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	testutil "github.com/patriciapersi/colabore-api/util"
	"github.com/stretchr/testify/assert"
)

func TestPostSolicitaFerias_(t *testing.T) {
	if err := testutil.LoadEnv(); err != nil {
		t.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	testCases := []struct {
		description  string
		setupBody    bool
		header       map[string]string
		expected     int
		expectedDesc string
	}{
		{
			description:  "Inclusão de Gestores",
			setupBody:    true,
			header:       config.SetupHeadersAgente(),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Inclusão de Gestores",
			setupBody:    false,
			header:       config.SetupHeadersAgente(),
			expected:     http.StatusBadRequest,
			expectedDesc: "Corpo da requisição não contém",
		},
		{
			description:  "Inclusão de Gestores",
			setupBody:    false,
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
				body = config.PostGestoresRequestBody()

			}

			resp, err := api.Client.R().
				SetHeaders(tc.header).
				SetBody(body).
				Post(api.EndpointsAgente["Gestor"])

			assert.NoError(t, err, "Erro ao fazer a requisição para %s", tc.description)
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado para %s", tc.description)
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}
}
