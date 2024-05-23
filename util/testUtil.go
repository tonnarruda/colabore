package testUtil

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func SetupEnv(envs map[string]string) {
	for key, value := range envs {
		os.Setenv(key, value)
	}
}

func LoadEnv() error {
	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load("c:\\workspace\\colabore-api\\.env")
	if err != nil {
		return fmt.Errorf("Erro ao carregar o arquivo .env: %v", err)
	}
	return nil
}
