package config

import (
	"os"

	"github.com/go-resty/resty/v2"
)

const BaseURL = "https://fortescolabore2.fortestecnologia.com.br/homolog-next/api"

func SetupClient() *resty.Client {
	return resty.New()
}

func SetupHeaders() map[string]string {
	return map[string]string{
		"Content-Type":                  "application/json",
		"x-api-key":                     os.Getenv("API_KEY"),
		"cnpj-licenciado":               os.Getenv("CNPJ_LICENCIADO"),
		"colabore-signature-expiration": os.Getenv("COLABORE_SIGNATURE_EXPIRATION"),
		"colabore-signature":            os.Getenv("COLABORE_SIGNATURE"),
	}
}
