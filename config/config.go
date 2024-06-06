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
		"colabore-signature-expiration": "3295537561",
		"colabore-signature":            "MEQCIHVn/5QZgqle6RyC2Lf8O/Eomw7RNvY7Em4GRWfawFj7AiBxrI0pdWE15EjfuKsQTlGYXaeXhN96EY3i34w6iUv/zA==",
	}
}
