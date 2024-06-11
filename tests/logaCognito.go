package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

const (
	USER_POOL_ID = "us-east-1_LjA7B5WOO" // Substitua pelo ID do seu User Pool
	CLIENT_ID    = "your_client_id"      // Substitua pelo ID do seu Cliente
	USERNAME     = "60515860409"         // Substitua pelo nome de usuário
	PASSWORD     = "12345678"            // Substitua pela senha do usuário
)

func main() {
	// Carregar a configuração padrão do SDK
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Criar um cliente CognitoIdentityProvider
	cognitoClient := cognitoidentityprovider.NewFromConfig(cfg)

	// Configurar os parâmetros da requisição
	params := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: "USER_PASSWORD_AUTH",
		AuthParameters: map[string]string{
			"USERNAME": USERNAME,
			"PASSWORD": PASSWORD,
		},
		ClientId: aws.String(CLIENT_ID),
	}

	// Chamar o método InitiateAuth
	authResponse, err := cognitoClient.InitiateAuth(context.TODO(), params)
	if err != nil {
		log.Fatalf("failed to authenticate user, %v", err)
	}

	// Imprimir o token retornado
	fmt.Printf("ID Token: %s\n", *authResponse.AuthenticationResult.IdToken)
	fmt.Printf("Access Token: %s\n", *authResponse.AuthenticationResult.AccessToken)
	fmt.Printf("Refresh Token: %s\n", *authResponse.AuthenticationResult.RefreshToken)
}
