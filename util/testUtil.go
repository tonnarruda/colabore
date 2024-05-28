package testUtil

import (
	"os"
)

func SetupEnv(envs map[string]string) {
	for key, value := range envs {
		os.Setenv(key, value)
	}
}
