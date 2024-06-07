package config

import (
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

type API struct {
	Client    *resty.Client
	Endpoints map[string]string
}

const BaseURL = "https://fortescolabore2.fortestecnologia.com.br/homolog-next/api"

func SetupClient() *resty.Client {
	return resty.New()
}

func SetupHeaders() map[string]string {
	godotenv.Load(".env")
	return map[string]string{
		"Content-Type":                  "application/json",
		"x-api-key":                     os.Getenv("X_API_KEY"),
		"cnpj-licenciado":               os.Getenv("CNPJ_LICENCIADO"),
		"colabore-signature-expiration": os.Getenv("COLABORE_SIGNATURE_EXPIRATION"),
		"colabore-signature":            os.Getenv("COLABORE_SIGNATURE"),
	}
}

func SetupApi() *API {

	api := resty.New().
		SetBaseURL("https://fortescolabore2.fortestecnologia.com.br/homolog-next")

	endpoints := map[string]string{
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

	return &API{
		Client:    api,
		Endpoints: endpoints,
	}

}
