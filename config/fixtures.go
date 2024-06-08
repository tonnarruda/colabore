package config

import (
	"time"

	"github.com/google/uuid"
)

var nrInsc string = "10821992"
var cpf string = "60515860409"

// NewRequestBody retorna o corpo da requisição formatado conforme necessário.
func DefinicoesRequestBody() map[string]interface{} {
	return map[string]interface{}{
		"Definicoes": []map[string]interface{}{
			{
				"NrInscEmpregador": nrInsc,
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
		"NrInscEmpregador": nrInsc,
		"MensagemTitulo":   "Teste automatizado",
		"MensagemCorpo":    "Mensagem enviada pelo teste automatizado",
		"DataMensagem":     time.Now().Format("02/01/2006"),
		"Colaboradores": []map[string]interface{}{
			{
				"CPF": cpf,
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
		"NrInscEmpregador": nrInsc,
		"ListaCPF":         []string{cpf},
	}
}

func PostInformacoesFeriasEmpregadoRequestBody() map[string]interface{} {
	return map[string]interface{}{
		"Colaboradores": []interface{}{
			map[string]interface{}{
				"CPF":                            cpf,
				"NrInscEmpregador":               "10821992",
				"Matricula":                      "000031",
				"SolicitouAdiantamento13":        false,
				"DiasDisponiveisAbonoPecuniario": 10,
				"DiasDisponiveis":                30,
				"InicioPeriodoConcessivo":        "2022-06-29",
				"FimPeriodoConcessivo":           "2023-07-30",
			},
		},
	}
}

func PostDispositivosRequestBody() map[string]interface{} {
	return map[string]interface{}{
		"Cnpj":          "63542443000124",
		"DispositivoId": "31e18fe4151b96cb",
		"Status":        1,
		"ListaEmpresas": []string{"10821992"},
	}
}

func PostSolicitaFeriasRequestBody() map[string]interface{} {
	return map[string]interface{}{
		"Ferias": []interface{}{
			map[string]interface{}{
				"CPF":                      cpf,
				"NrInscEmpregador":         nrInsc,
				"Matricula":                "000031",
				"SolicitouAdiantamento13":  true,
				"SolicitouAbonoPecuniario": true,
				"StatusSolicitacao":        4,
				"InicioPeriodoGozo":        "2024-06-29",
				"FimPeriodoGozo":           "2024-07-08",
			},
		},
	}
}
