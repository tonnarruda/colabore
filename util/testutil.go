// Dentro do seu pacote testutil

package testutil

import "os"

// SetupEnv configura as variáveis de ambiente seletivamente
func SetupEnv(envs map[string]string) {
	for key, value := range envs {
		os.Setenv(key, value)
	}
}
