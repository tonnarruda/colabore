package config

import (
	"os"

	"github.com/go-resty/resty/v2"
	testutil "github.com/patriciapersi/colabore-api/util"
)

type API struct {
	Client          *resty.Client
	EndpointsAgente map[string]string
	EndpointsApp    map[string]string
}

func SetupHeadersAgente() map[string]string {
	testutil.LoadEnv()
	return map[string]string{
		"Content-Type":                  "application/json",
		"x-api-key":                     os.Getenv("X_API_KEY"),
		"cnpj-licenciado":               os.Getenv("CNPJ_LICENCIADO"),
		"colabore-signature-expiration": os.Getenv("COLABORE_SIGNATURE_EXPIRATION"),
		"colabore-signature":            os.Getenv("COLABORE_SIGNATURE"),
	}
}

func SetupHeadersApp() map[string]string {
	testutil.LoadEnv()
	return map[string]string{
		"Content-Type": "application/json",
		"awsauthtoken": "eyJraWQiOiJVWEkwNXFBb2VDeTJPY1BGM21iOGJtWXlyWGk4N1N3K1ZqWXhIRWV3K2RVPSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiIyYjljZmU1Mi1hZGE0LTQ0NzMtYmRkYi05NDMyOTRkNzc0OTciLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiaXNzIjoiaHR0cHM6XC9cL2NvZ25pdG8taWRwLnVzLWVhc3QtMS5hbWF6b25hd3MuY29tXC91cy1lYXN0LTFfTGpBN0I1V09PIiwiY29nbml0bzp1c2VybmFtZSI6IjI2Mzg0ODEzNzg1IiwiYXVkIjoiNDVxb2RmNWhsNGdncnY1b2tpOGM5amdnYmoiLCJldmVudF9pZCI6IjZkNmJhZmE4LTdmNDQtNDc3My05NDFjLTgzZTQ5MWE1NzAwZSIsInRva2VuX3VzZSI6ImlkIiwiYXV0aF90aW1lIjoxNzAwMjQxNjM2LCJleHAiOjE3MDAyNDUyMzYsImlhdCI6MTcwMDI0MTYzNiwiZmFtaWx5X25hbWUiOiJNYXJpYSIsImVtYWlsIjoiZm9ydGVzY29sYWJvcmVAZm9ydGVzdGVjbm9sb2dpYS5jb20uYnIifQ.CjCHh0Izg4Tvd9ykrHDknikJfU5P2yiB241ijqjo1h0atv_bfZDpKld6ymiHm6gXqQqRBhd0V2hmpDiCSJgPKEDWAih8zz1s2cND7P3HoQzuR2-1C5B3ZhU1SjX3R7K4Ui2L-sy6skkw2ZWb0s9gDlpdsLyoR2S6UefWaAumt1dh6QqHFw4zKGztl-TS9fGwMEFKi7nF-DRl-PbVKE6ALPXsf0MGwm947lSNpZU04zTzfJHHSh-ARIvZrQMSNb-3AjL2z9PAAA8On-KJJUGq2Ztiy9HO3m9a7cy75GjwQa4jjk_480IkNKP7923K3ugby5SeXw_8x0AL7VODQm_UVQ",
	}
}

func SetupApi() *API {

	api := resty.New().
		SetBaseURL("https://fortescolabore2.fortestecnologia.com.br/homolog-next")

	endpointsAgente := map[string]string{
		"ImagemReferencia":                         "/api/agente/Arquivo/ImagemReferencia",
		"ArquivoRelatorioPontoURLTemporaria":       "/api/agente/Arquivo/RelatorioPonto/URLTemporaria",
		"AssinaturaLiberar":                        "/api/agente/Assinatura/Liberar",
		"Assinatura":                               "/api/agente/Assinatura",
		"AssinaturaContestacao":                    "/api/agente/Assinatura/Contestacao",
		"AssinaturaContestacaoStatus":              "/api/agente/Assinatura/Contestacao/Status",
		"AssinaturaRelatorioPonto":                 "/api/agente/Assinatura/RelatorioPonto",
		"AssinaturaDigital":                        "/api/agente/AssinaturaDigital",
		"AssinaturaDigitalFormData":                "/api/agente/AssinaturaDigital/FormData",
		"Candidato":                                "/api/agente/Candidato",
		"CandidatoRetificar":                       "/api/agente/Candidato/Retificar",
		"Acesso":                                   "/api/agente-acesso", //NAO SERÁ AUTOMATIZADO
		"Agente":                                   "/api/agente",        //NAO SERÁ AUTOMATIZADO
		"Colaborador":                              "/api/agente/Colaborador",
		"ColaboradorAtivacao":                      "/api/agente/Colaborador/Ativacao",
		"ColaboradorFreemium":                      "/api/agente/Colaborador/Freemium",
		"ColaboradorPreemium":                      "/api/agente/Colaborador/Premium",
		"ColaboradorEnderecoConfirmacao":           "/api/agente/Colaborador/Endereco/Confirmacao",
		"ColaboradorEnderecos":                     "/api/agente/Colaborador/Enderecos",
		"ColaboradorDefinicoes":                    "/api/agente/Colaborador/Definicoes",
		"Declaracao":                               "/api/agente/Declaracao",
		"Dispositivos":                             "/api/agente/Dispositivos",
		"DispositivosStatus":                       "/api/agente/Dispositivos/Status",
		"FeriasInformacoes":                        "/api/agente/Ferias/Informacoes",
		"Ferias":                                   "/api/agente/Ferias",
		"FolhaDePagamento":                         "/api/agente/FolhaDePagamento",
		"Gestor":                                   "/api/agente/Gestor",
		"GestorRH":                                 "/api/agente/Gestor/RH",
		"GETimagem":                                "/api/agente/Imagem",
		"LicencaReconhecimentoFacial":              "/api/agente/Licenca/ReconhecimentoFacial",
		"LicenciadoEmailParaNotificacoesDaEmpresa": "/api/agente/Licenciado/EmailParaNotificacoesDaEmpresa",
		"LicenciadoLogo":                           "/api/agente/Licenciado/Logo",
		"LicenciadoDefinicoes":                     "/api/agente/Licenciado/Definicoes",
		"Mensagem":                                 "/api/agente/Mensagem",
		"Numerador":                                "/api/agente/Numerador",
		"Pesquisa":                                 "/api/agente/Pesquisa",
		"PesquisaRespostas":                        "/api/agente/Pesquisa/Respostas",
		"PesquisaRespostasPorPagina":               "/api/agente/Pesquisa/RespostasPorPagina",
		"PesquisaResposta":                         "/api/agente/Pesquisa/Resposta",
		"PontoBatidas":                             "/api/agente/Ponto/Batidas",
		"PontoProcessamentoRedefinicao":            "/api/agente/Ponto/Processamento/Redefinicao",
		"PontoAssinaturaRetroativaSolicitacao":     "/api/agente/Ponto/AssinaturaRetroativa/Solicitacao",
		"PontoObterTimeZone":                       "/api/agente/Ponto/ObterTimeZone",
		"Termo":                                    "/api/agente/Termo",
		"VersaoApp":                                "/api/agente/VersaoApp",
	}

	endpointsApp := map[string]string{
		"Appferias": "/api/app/Ferias",
	}

	return &API{
		Client:          api,
		EndpointsAgente: endpointsAgente,
		EndpointsApp:    endpointsApp,
	}

}
