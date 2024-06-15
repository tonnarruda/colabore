package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	"github.com/stretchr/testify/assert"
)

func TestPostSolicitaAbono(t *testing.T) {

	testCases := []struct {
		description  string
		header       map[string]string
		setupBody    bool
		expected     int
		expectedDesc string
	}{
		{
			description:  "Teste Solicita abono",
			header:       config.SetupHeadersAgente(),
			setupBody:    true,
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Teste envio de solicitações sem body",
			header:       config.SetupHeadersAgente(),
			setupBody:    false,
			expected:     http.StatusBadRequest,
			expectedDesc: "ERRO",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			api := config.SetupApi()

			// Configura os parâmetros do corpo da requisição se necessário
			var body interface{}
			if tc.setupBody {
				body = config.PostSolicitaAbonoBody()
			}

			resp, err := api.Client.R().
				SetHeaders(tc.header).
				SetBody(body).
				Post(api.EndpointsAgente["AbonoAprovaReprova"])

			assert.NoError(t, err, "Erro ao fazer a requisição para %s", tc.description)
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado para %s", tc.description)
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}

}
