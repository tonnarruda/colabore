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
		expected         int
		expectedDesc     string
	}{
		{
			description:      "Buscar Imagem com sucesso",
			cpf:              "60515860409",
			NrInscEmpregador: "10821992",
			expected:         http.StatusOK,
			expectedDesc:     "Sucesso",
		},
		{
			description:      "Buscar Imagem com NrInsc Invalido",
			cpf:              "60515860409",
			NrInscEmpregador: "00000000",
			expected:         http.StatusBadRequest,
			expectedDesc:     "Arquivo não encontrado",
		},
		{
			description:      "Buscar Imagem com NrInsc Vazio",
			cpf:              "60515860409",
			NrInscEmpregador: "",
			expected:         http.StatusBadRequest,
			expectedDesc:     "CaminhoArquivo",
		},
		{
			description:      "Buscar Imagem com CPF Vazio",
			cpf:              "",
			NrInscEmpregador: "10821992",
			expected:         http.StatusBadRequest,
			expectedDesc:     "CaminhoArquivo",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			client := config.SetupClient()
			url := config.BaseURL + "/agente/Arquivo/ImagemReferencia"

			queryParams := map[string]string{
				"CPF":              tc.cpf,
				"NrInscEmpregador": tc.NrInscEmpregador,
			}

			resp, err := client.R().
				SetHeaders(config.SetupHeaders()).
				SetQueryParams(queryParams).
				Get(url)

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}
}