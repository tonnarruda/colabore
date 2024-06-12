package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	testutil "github.com/patriciapersi/colabore-api/util"
	"github.com/stretchr/testify/assert"
)

func TestGetLicenca(t *testing.T) {
	testutil.LoadEnv()

	testCases := []struct {
		description  string
		header       map[string]string
		expected     int
		expectedDesc string
	}{
		{
			description:  "Busca de Reconhecimento Facial com Sucesso",
			header:       config.SetupHeadersAgente(),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Busca de Reconhecimento Facial com Sucesso",
			header:       map[string]string{},
			expected:     http.StatusUnauthorized,
			expectedDesc: "Unauthorized",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			api := config.SetupApi()
			resp, err := api.Client.R().
				SetHeaders(tc.header).
				Get(api.EndpointsAgente["LicencaReconhecimentoFacial"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}
}
