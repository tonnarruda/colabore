package config

import (
	"github.com/go-resty/resty/v2"
)

const BaseURL = "https://fortescolabore2.fortestecnologia.com.br/homolog-next/api"

func SetupClient() *resty.Client {
	return resty.New()
}

func SetupHeaders() map[string]string {

	return map[string]string{
		"Content-Type":                  "application/json",
		"x-api-key":                     "ACTIjCHfsj1QZulkDK9oZ5bfPrlparVS44cb9o7F",
		"cnpj-licenciado":               "63542443000124",
		"colabore-signature-expiration": "1709743637",
		"colabore-signature":            "MEQCIHzCN4gb+VSQqTmYLnHCjo8pQASiuQ4Nv+B1Jd9jwFtWAiBO0ZdfSciwKOTkPf6LIJHMtyHSGpLzLXgFXH1znVxrVA==",
	}
}
