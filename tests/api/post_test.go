package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestNetHTTP(t *testing.T) {
	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		t.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Obtém os valores das variáveis de ambiente
	cnpjLicenciado := os.Getenv("CNPJ_LICENCIADO")
	signatureExpiration := os.Getenv("COLABORE_SIGNATURE_EXPIRATION")
	signature := os.Getenv("COLABORE_SIGNATURE")

	// Define o corpo da requisição
	body := map[string]interface{}{
		"Definicoes": []map[string]interface{}{
			{
				"NrInscEmpregador": "10821992",
				"Ferias": map[string]interface{}{
					"AntecedenciaMinima":     20,
					"HabilitaFerias":         true,
					"ExigeAprovacaoDoGestor": true,
				},
			},
		},
	}

	// Converte o corpo da requisição para JSON
	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Erro ao codificar JSON: %v", err)
	}

	// Cria uma requisição POST
	req, err := http.NewRequest("POST", "https://fortescolabore2.fortestecnologia.com.br/homolog-next/api/agente/licenciado/definicoes", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Erro ao criar a requisição POST: %v", err)
	}

	// Define os cabeçalhos da requisição
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", "ACTIjCHfsj1QZulkDK9oZ5bfPrlparVS44cb9o7F")
	req.Header.Set("cnpj-licenciado", cnpjLicenciado)
	req.Header.Set("colabore-signature-expiration", signatureExpiration)
	req.Header.Set("colabore-signature", signature)

	// Faz a requisição HTTP
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("Erro ao fazer a requisição HTTP: %v", err)
	}
	defer res.Body.Close()

	// Verifica o código de status da resposta
	if res.StatusCode != http.StatusOK {
		t.Errorf("Código de status inesperado: esperado %d, recebido %d", http.StatusOK, res.StatusCode)
	}

	// Processa o corpo da resposta
	var response map[string]string
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Fatalf("Erro ao decodificar JSON da resposta: %v", err)
	}
}
