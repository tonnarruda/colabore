package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	"github.com/stretchr/testify/assert"
)

func TestGetImagem(t *testing.T) {

	testCases := []struct {
		description      string
		cpf              string
		NrInscEmpregador string
		header           map[string]string
		expected         int
		expectedDesc     string
	}{
		{
			description:      "Buscar imagem com sucesso",
			cpf:              "60515860409",
			NrInscEmpregador: "10821992",
			header:           config.SetupHeadersAgente(),
			expected:         http.StatusOK,
			expectedDesc:     "Sucesso",
		},
		{
			description:      "Buscar imagem com NrInscEmpregador inválido",
			cpf:              "60515860409",
			NrInscEmpregador: "00000000",
			header:           config.SetupHeadersAgente(),
			expected:         http.StatusBadRequest,
			expectedDesc:     "Arquivo não encontrado",
		},
		{
			description:      "Buscar imagem com NrInscEmpregador vazio",
			cpf:              "60515860409",
			NrInscEmpregador: "",
			header:           config.SetupHeadersAgente(),
			expected:         http.StatusBadRequest,
			expectedDesc:     "CaminhoArquivo",
		},
		{
			description:      "Buscar imagem com CPF vazio",
			cpf:              "",
			NrInscEmpregador: "10821992",
			header:           config.SetupHeadersAgente(),
			expected:         http.StatusBadRequest,
			expectedDesc:     "CaminhoArquivo",
		},
		{
			description:      "Buscar imagem sem cabeçalhos de autenticação",
			cpf:              "60515860409",
			NrInscEmpregador: "10821992",
			header:           map[string]string{},
			expected:         http.StatusUnauthorized,
			expectedDesc:     "Unauthorized",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			api := config.SetupApi()
			queryParams := map[string]string{
				"CPF":              tc.cpf,
				"NrInscEmpregador": tc.NrInscEmpregador,
			}

			resp, err := api.Client.R().
				SetHeaders(tc.header).
				SetQueryParams(queryParams).
				Get(api.EndpointsAgente["ImagemReferencia"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}
}
