package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	testutil "github.com/patriciapersi/colabore-api/util"
	"github.com/stretchr/testify/assert"
)

func TestGetTimeZone(t *testing.T) {
	testutil.LoadEnv()

	// Define os casos de teste em uma tabela
	testCases := []struct {
		description  string
		latitude     string
		longitude    string
		header       map[string]string
		expected     int
		expectedDesc string
	}{
		{
			description:  "Envio de latitude e longitude válidas",
			latitude:     "-22.9519",
			longitude:    "-43.2105",
			header:       config.SetupHeadersAgente(),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Envio de latitude e longitude inválidas",
			latitude:     "latitude",
			longitude:    "longitude",
			header:       config.SetupHeadersAgente(),
			expected:     http.StatusBadRequest,
			expectedDesc: "latitude,longitude",
		},
		{
			description:  "Envio de latitude e longitude vazias",
			latitude:     "",
			longitude:    "",
			header:       config.SetupHeadersAgente(),
			expected:     http.StatusBadRequest,
			expectedDesc: "latitude,longitude",
		},
		{
			description:  "Envio de latitude e longitude válidas - header vazio",
			latitude:     "-22.9519",
			longitude:    "-43.2105",
			header:       map[string]string{},
			expected:     http.StatusUnauthorized,
			expectedDesc: "Unauthorized",
		},
	}

	// Itera sobre os casos de teste
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			api := config.SetupApi()
			queryParams := map[string]string{
				"latitude":  tc.latitude,
				"longitude": tc.longitude,
			}

			resp, err := api.Client.R().
				SetHeaders(tc.header).
				SetQueryParams(queryParams).
				Get(api.EndpointsAgente["PontoObterTimeZone"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}
}
