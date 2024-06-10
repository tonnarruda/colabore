package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	"github.com/stretchr/testify/assert"

	testutil "github.com/patriciapersi/colabore-api/util"
)

func TestGetImagem(t *testing.T) {
	if err := testutil.LoadEnv(); err != nil {
		t.Fatalf("Erro ao carregar o arquivo .env: %v", err)
		t.Fatalf("%v", err)
	}

	testCases := []struct {
		description      string
		cpf              string
		NrInscEmpregador string
		header           map[string]string
		expected         int
		expectedDesc     string
	}{
		{
			description:      "Buscar Imagem com sucesso",
			cpf:              "60515860409",
			NrInscEmpregador: "10821992",
			header:           config.SetupHeadersAgente(),
			expected:         http.StatusOK,
			expectedDesc:     "Sucesso",
		},
		{
			description:      "Buscar Imagem com NrInsc Invalido",
			cpf:              "60515860409",
			NrInscEmpregador: "00000000",
			header:           config.SetupHeadersAgente(),
			expected:         http.StatusBadRequest,
			expectedDesc:     "Arquivo não encontrado",
		},
		{
			description:      "Buscar Imagem com NrInsc Vazio",
			cpf:              "60515860409",
			NrInscEmpregador: "",
			header:           config.SetupHeadersAgente(),
			expected:         http.StatusBadRequest,
			expectedDesc:     "CaminhoArquivo",
		},
		{
			description:      "Buscar Imagem com CPF Vazio",
			cpf:              "",
			NrInscEmpregador: "10821992",
			header:           config.SetupHeadersAgente(),
			expected:         http.StatusBadRequest,
			expectedDesc:     "CaminhoArquivo",
		},
		{
			description:      "Buscar Imagem com sucesso",
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
