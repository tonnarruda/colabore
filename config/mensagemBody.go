package config

// NewRequestBody retorna o corpo da requisição formatado conforme necessário.

func MensagensRequestBody() map[string]interface{} {
	return map[string]interface{}{
		"Mensagens": []map[string]interface{}{
			{
				"ID":               "236",
				"TpInscEmpregador": "1",
				"NrInscEmpregador": "10821992",
				"MensagemTitulo":   "Feriado new",
				"MensagemCorpo":    "Saiba o que funciona no feriado de maio em nossa empresa",
				"DataMensagem":     "22/05/2024",
				"Colaboradores": []map[string]interface{}{
					{
						"CPF": "60515860409",
						"Contrato": map[string]interface{}{
							"Matricula": "000031",
							"Cargo":     "ALMOXARIFE",
						},
					},
				},
			},
		},
	}
}
