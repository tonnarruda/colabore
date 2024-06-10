package config

import (
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

type API struct {
	Client          *resty.Client
	EndpointsAgente map[string]string
	EndpointsApp    map[string]string
}

const BaseURL = "https://fortescolabore2.fortestecnologia.com.br/homolog-next/api"

func SetupClient() *resty.Client {
	return resty.New()
}

func SetupHeadersAgente() map[string]string {
	godotenv.Load(".env")
	return map[string]string{
		"Content-Type":                  "application/json",
		"x-api-key":                     os.Getenv("X_API_KEY"),
		"cnpj-licenciado":               os.Getenv("CNPJ_LICENCIADO"),
		"colabore-signature-expiration": os.Getenv("COLABORE_SIGNATURE_EXPIRATION"),
		"colabore-signature":            os.Getenv("COLABORE_SIGNATURE"),
	}
}

func SetupHeadersApp() map[string]string {
	godotenv.Load(".env")
	return map[string]string{
		"Content-Type": "application/json",
		"awsauthtoken": "eyJraWQiOiJVWEkwNXFBb2VDeTJPY1BGM21iOGJtWXlyWGk4N1N3K1ZqWXhIRWV3K2RVPSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiIyYjljZmU1Mi1hZGE0LTQ0NzMtYmRkYi05NDMyOTRkNzc0OTciLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiaXNzIjoiaHR0cHM6XC9cL2NvZ25pdG8taWRwLnVzLWVhc3QtMS5hbWF6b25hd3MuY29tXC91cy1lYXN0LTFfTGpBN0I1V09PIiwiY29nbml0bzp1c2VybmFtZSI6IjI2Mzg0ODEzNzg1IiwiYXVkIjoiNDVxb2RmNWhsNGdncnY1b2tpOGM5amdnYmoiLCJldmVudF9pZCI6IjZkNmJhZmE4LTdmNDQtNDc3My05NDFjLTgzZTQ5MWE1NzAwZSIsInRva2VuX3VzZSI6ImlkIiwiYXV0aF90aW1lIjoxNzAwMjQxNjM2LCJleHAiOjE3MDAyNDUyMzYsImlhdCI6MTcwMDI0MTYzNiwiZmFtaWx5X25hbWUiOiJNYXJpYSIsImVtYWlsIjoiZm9ydGVzY29sYWJvcmVAZm9ydGVzdGVjbm9sb2dpYS5jb20uYnIifQ.CjCHh0Izg4Tvd9ykrHDknikJfU5P2yiB241ijqjo1h0atv_bfZDpKld6ymiHm6gXqQqRBhd0V2hmpDiCSJgPKEDWAih8zz1s2cND7P3HoQzuR2-1C5B3ZhU1SjX3R7K4Ui2L-sy6skkw2ZWb0s9gDlpdsLyoR2S6UefWaAumt1dh6QqHFw4zKGztl-TS9fGwMEFKi7nF-DRl-PbVKE6ALPXsf0MGwm947lSNpZU04zTzfJHHSh-ARIvZrQMSNb-3AjL2z9PAAA8On-KJJUGq2Ztiy9HO3m9a7cy75GjwQa4jjk_480IkNKP7923K3ugby5SeXw_8x0AL7VODQm_UVQ",
	}
}

func SetupApi() *API {

	api := resty.New().
		SetBaseURL("https://fortescolabore2.fortestecnologia.com.br/homolog-next")

	endpointsAgente := map[string]string{
		"GETarquivoImagemReferencia":                   "/api/agente/Arquivo/ImagemReferencia",
		"DELETEarquivoImagemReferencia":                "/api/agente/Arquivo/ImagemReferencia",
		"GETarquivoRelatorioPontoURLTemporaria":        "/api/agente/Arquivo/RelatorioPonto/URLTemporaria",
		"POSTassinaturaLiberar":                        "/api/agente/Assinatura/Liberar",
		"GETassinatura":                                "/api/agente/Assinatura",
		"DELETEassinatura":                             "/api/agente/Assinatura",
		"GETassinaturaContestacao":                     "/api/agente/Assinatura/Contestacao",
		"POSTassinaturaContestacaoStatus":              "/api/agente/Assinatura/Contestacao/Status",
		"POSTassinaturaRelatorioPonto":                 "/api/agente/Assinatura/RelatorioPonto",
		"POSTassinaturaDigital":                        "/api/agente/AssinaturaDigital",
		"POSTassinaturaDigitalFormData":                "/api/agente/AssinaturaDigital/FormData",
		"POSTcandidato":                                "/api/agente/Candidato",
		"GETcandidato":                                 "/api/agente/Candidato",
		"PUTcandidatoRetificar":                        "/api/agente/Candidato/Retificar",
		"POSTacesso":                                   "/api/agente-acesso",
		"POSTagente":                                   "/api/agente",
		"POSTcolaborador":                              "/api/agente/Colaborador",
		"DELETEcolaborador":                            "/api/agente/Colaborador",
		"PUTcolaboradorAtivacao":                       "/api/agente/Colaborador/Ativacao",
		"GETcolaboradorFreemium":                       "/api/agente/Colaborador/Freemium",
		"GETcolaboradorPremium":                        "/api/agente/Colaborador/Premium",
		"PUTcolaboradorEnderecoConfirmacao":            "/api/agente/Colaborador/Endereco/Confirmacao",
		"POSTcolaboradorEnderecos":                     "/api/agente/Colaborador/Enderecos",
		"PUTcolaboradorDefinicoes":                     "/api/agente/Colaborador/Definicoes",
		"DELETEdeclaracao":                             "/api/agente/Declaracao",
		"POSTdeclaracao":                               "/api/agente/Declaracao",
		"GETdispositivos":                              "/api/agente/Dispositivos",
		"DELETEdispositivos":                           "/api/agente/Dispositivos",
		"POSTdispositivosStatus":                       "/api/agente/Dispositivos/Status",
		"POSTferiasInformacoes":                        "/api/agente/Ferias/Informacoes",
		"GETferias":                                    "/api/agente/Ferias",
		"POSTferias":                                   "/api/agente/Ferias",
		"POSTfolhaDePagamento":                         "/api/agente/FolhaDePagamento",
		"DELETEfolhaDePagamento":                       "/api/agente/FolhaDePagamento",
		"POSTgestor":                                   "/api/agente/Gestor",
		"DELETEgestor":                                 "/api/agente/Gestor",
		"POSTgestorRH":                                 "/api/agente/Gestor/RH",
		"GETimagem":                                    "/api/agente/Imagem",
		"GETlicencaReconhecimentoFacial":               "/api/agente/Licenca/ReconhecimentoFacial",
		"POSTlicenciadoEmailParaNotificacoesDaEmpresa": "/api/agente/Licenciado/EmailParaNotificacoesDaEmpresa",
		"POSTlicenciadoLogo":                           "/api/agente/Licenciado/Logo",
		"POSTlicenciadoDefinicoes":                     "/api/agente/Licenciado/Definicoes",
		"POSTmensagem":                                 "/api/agente/Mensagem",
		"DELETEmensagem":                               "/api/agente/Mensagem",
		"PUTnumerador":                                 "/api/agente/Numerador",
		"POSTpesquisa":                                 "/api/agente/Pesquisa",
		"DELETEpesquisa":                               "/api/agente/Pesquisa",
		"GETpesquisaRespostas":                         "/api/agente/Pesquisa/Respostas",
		"GETpesquisaRespostasPorPagina":                "/api/agente/Pesquisa/RespostasPorPagina",
		"POSTpesquisaResposta":                         "/api/agente/Pesquisa/Resposta",
		"DELETEpesquisaResposta":                       "/api/agente/Pesquisa/Resposta",
		"POSTpontoBatidas":                             "/api/agente/Ponto/Batidas",
		"POSTpontoProcessamentoRedefinicao":            "/api/agente/Ponto/Processamento/Redefinicao",
		"POSTpontoAssinaturaRetroativaSolicitacao":     "/api/agente/Ponto/AssinaturaRetroativa/Solicitacao",
		"GETpontoObterTimeZone":                        "/api/agente/Ponto/ObterTimeZone",
		"PUTtermo":                                     "/api/agente/Termo",
		"PUTversaoApp":                                 "/api/agente/VersaoApp",
	}

	endpointsApp := map[string]string{
		"POSTappferias": "/api/app/Ferias",
	}

	return &API{
		Client:          api,
		EndpointsAgente: endpointsAgente,
		EndpointsApp:    endpointsApp,
	}

}
