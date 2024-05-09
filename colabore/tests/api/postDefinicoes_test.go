package main

import (
	"os"
	"testing"

	"github.com/h2non/baloo"
	"github.com/joho/godotenv"
)

var test = baloo.New("https://fortescolabore2.fortestecnologia.com.br")

func TestPostColabore(t *testing.T) {
	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		t.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Obtém os valores das variáveis de ambiente
	cnpjLicenciado := os.Getenv("CNPJ_LICENCIADO")
	signatureExpiration := os.Getenv("COLABORE_SIGNATURE_EXPIRATION")
	signature := os.Getenv("COLABORE_SIGNATURE")

	// Define o corpo da requisição com o JSON fornecido
	body := map[string]interface{}{
		"Definicoes": []map[string]interface{}{
			{
				"NrInscEmpregador": "10821992",
				"Ferias": map[string]interface{}{
					"AntecedenciaMinima":     15,
					"HabilitaFerias":         true,
					"ExigeAprovacaoDoGestor": true,
				},
			},
		},
	}

	// Define os cabeçalhos da requisição
	headers := map[string]string{
		"Content-Type":                  "application/json",
		"x-api-key":                     "ACTIjCHfsj1QZulkDK9oZ5bfPrlparVS44cb9o7F",
		"cnpj-licenciado":               cnpjLicenciado,
		"colabore-signature-expiration": signatureExpiration,
		"colabore-signature":            signature,
	}

	// Faz a requisição POST e verifica a resposta
	test.Post("/homolog-next/api/agente/licenciado/definicoes").
		JSON(body).
		SetHeaders(headers).
		Expect(t).
		Status(200).
		Header("Server", "apache").
		Type("json").
		JSON(map[string]string{"bar": "foo"}).
		Done()
}
