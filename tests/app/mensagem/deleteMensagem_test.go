package main

import (
	"log"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/patriciapersi/colabore-api/config"
	"github.com/stretchr/testify/assert"
)

func getMessageID() string {
	api := config.SetupApi()
	requestBody := config.MensagensRequestBody()
	id := requestBody["ID"].(string)
	resp, _ := api.Client.R().
		SetHeaders(config.SetupHeadersAgente()).
		SetBody(requestBody).
		Post(api.EndpointsAgente["Mensagem"])

	if resp.StatusCode() != http.StatusOK {
		log.Printf("Unexpected status code: %d", resp.StatusCode())
	}

	return id
}

func TestDeleteMensagens(t *testing.T) {

	testsCases := []struct {
		description  string
		header       map[string]string
		id           string
		expected     int
		expectedDesc string
	}{
		{
			description:  "Mensagem existente",
			header:       config.SetupHeadersApp(),
			id:           getMessageID(),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "ID inexistente",
			header:       config.SetupHeadersApp(),
			id:           uuid.New().String(),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "ID vazio",
			header:       config.SetupHeadersApp(),
			id:           "",
			expected:     http.StatusBadRequest,
			expectedDesc: "MensagemId",
		},
		{
			description:  "Unauthorized",
			header:       map[string]string{},
			id:           "",
			expected:     http.StatusUnauthorized,
			expectedDesc: "Unauthorized",
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.description, func(t *testing.T) {

			api := config.SetupApi()

			resp, err := api.Client.R().
				SetHeaders(tc.header).
				SetBody(config.DeleteMensagemAppRequestBody(tc.id)).
				Delete(api.EndpointsApp["Mensagem"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")

		})
	}
}
