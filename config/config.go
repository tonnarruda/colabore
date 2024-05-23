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

// Captura o estado original das variáveis de ambiente especificadas
func CaptureOriginalEnv(headers []string) map[string]string {
	originalEnv := make(map[string]string)
	for _, key := range headers {
		originalEnv[key] = os.Getenv(key)
	}
	return originalEnv
}

// Restaura as variáveis de ambiente ao estado original
func RestoreEnv(originalEnv map[string]string) {
	for key, value := range originalEnv {
		os.Setenv(key, value)
	}
}

// Função modificada para capturar apenas as chaves das variáveis de ambiente
func GetEnvKeysFromHeaders() []string {
	headers := SetupHeaders()
	keys := make([]string, 0, len(headers))
	for key := range headers {
		keys = append(keys, key)
	}
	return keys
}
