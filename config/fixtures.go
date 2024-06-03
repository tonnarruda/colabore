package config

import (
	"time"

	"github.com/google/uuid"
)

// NewRequestBody retorna o corpo da requisição formatado conforme necessário.
func DefinicoesRequestBody() map[string]interface{} {
	return map[string]interface{}{
		"Definicoes": []map[string]interface{}{
			{
				"NrInscEmpregador": "10821992",
				"Ferias": map[string]interface{}{
					"AntecedenciaMinima":     15,
					"HabilitaFerias":         true,
					"ExigeAprovacaoDoGestor": true,
				},
			},
		},
	}
}

func MensagensRequestBody() map[string]interface{} {
	return map[string]interface{}{
		"ID":               uuid.New().String(),
		"TpInscEmpregador": "1",
		"NrInscEmpregador": "10821992",
		"MensagemTitulo":   "Teste automatizado",
		"MensagemCorpo":    "Mensagem enviada pelo teste automatizado",
		"DataMensagem":     time.Now().Format("02/01/2006"),
		"Colaboradores": []map[string]interface{}{
			{
				"CPF": "60515860409",
				"Contrato": map[string]interface{}{
					"Matricula": "000031",
					"Cargo":     "ALMOXARIFE",
				},
			},
		},
	}
}

func DeleteMensagensRequestBody(mensagemID string) map[string]interface{} {
	return map[string]interface{}{
		"MensagemId":       mensagemID,
		"NrInscEmpregador": "10821992",
		"ListaCPF":         []string{"60515860409"},
	}
}
