package testUtil

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func SetupEnv(envs map[string]string) {
	for key, value := range envs {
		os.Setenv(key, value)
	}
}

func LoadEnv() error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("erro ao obter o diretório de trabalho atual: %v", err)
	}

	// Vai subindo nos diretórios até encontrar o .env
	for {
		envPath := filepath.Join(cwd, ".env")
		if _, err := os.Stat(envPath); err == nil {
			// Carrega as variáveis de ambiente do arquivo .env
			err := godotenv.Load(envPath)
			if err != nil {
				return fmt.Errorf("erro ao carregar o arquivo .env: %v", err)
			}
			return nil
		}

		// Se não encontrou o arquivo .env, tenta o diretório pai
		parentDir := filepath.Dir(cwd)
		if parentDir == cwd {
			return fmt.Errorf("arquivo .env não encontrado")
		}
		cwd = parentDir
	}
}
