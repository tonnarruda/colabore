package main

import (
	"net/http"
	"testing"
	"time"

	"github.com/patriciapersi/colabore-api/config"
	testutil "github.com/patriciapersi/colabore-api/util"
	"github.com/stretchr/testify/assert"
)

func TestGetAssinaturaContestacao(t *testing.T) {
	testutil.LoadEnv()

	testCases := []struct {
		description  string
		params       map[string]string
		header       map[string]string
		expected     int
		expectedDesc string
	}{
		{
			description: "Buscar assinatura Contestada com sucesso",
			params: map[string]string{
				"UltimaAtualizacao": "0",
				"CPF":               "60515860409",
				"NrInscEmpregador":  "10821992",
				"Matricula":         "000031",
				"AnoMes":            time.Now().AddDate(0, 2, 0).Format("200601"),
			},
			header:       config.SetupHeadersAgente(),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description: "Buscar assinatura Contestada com parâmetros incompletos",
			params: map[string]string{
				"AnoMes": time.Now().AddDate(0, 2, 0).Format("012006"),
			},
			header:       config.SetupHeadersAgente(),
			expected:     http.StatusBadRequest,
			expectedDesc: "NrInscEmpregador",
		},
		{
			description: "Buscar assinatura sem cabeçalhos de autenticação",
			params: map[string]string{
				"AnoMes": time.Now().AddDate(0, 2, 0).Format("012006"),
			},
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
				SetQueryParams(tc.params).
				Get(api.EndpointsAgente["AssinaturaContestacao"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}
}
