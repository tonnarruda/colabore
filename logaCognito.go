package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	testutil "github.com/patriciapersi/colabore-api/util"
)

func main() {
	testutil.LoadEnv()

	// Definir a região da AWS que você deseja usar
	region := "us-west-1" // Substitua com a região desejada

	// Carregar a configuração do SDK com a região especificada
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Criar um cliente CognitoIdentityProvider
	cognitoClient := cognitoidentityprovider.NewFromConfig(cfg)

	clientDetails, err := json.MarshalIndent(cognitoClient, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal cognitoClient, %v", err)
	}
	fmt.Printf("Cognito Client Details: %s\n", string(clientDetails))

	// Configurar os parâmetros da requisição
	params := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: "USER_PASSWORD_AUTH",
		AuthParameters: map[string]string{
			"USERNAME": os.Getenv("USERNAME"),
			"PASSWORD": os.Getenv("PASSWORD"),
		},
		ClientId: aws.String(os.Getenv("CLIENT_ID")),
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
