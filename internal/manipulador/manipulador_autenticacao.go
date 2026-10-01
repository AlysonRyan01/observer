package manipulador

import (
	"go-api/internal/configuracao"
	"go-api/internal/requisicao"
	"go-api/internal/servico"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ManipuladorAutenticacao struct {
	servico *servico.ServicoAutenticacao
}

func NovoManipuladorAutenticacao(s *servico.ServicoAutenticacao) *ManipuladorAutenticacao {
	return &ManipuladorAutenticacao{servico: s}
}

func (m *ManipuladorAutenticacao) ManipularRegistro(c *gin.Context) {
	var req requisicao.RequisicaoUsuario

	if err := c.ShouldBindJSON(&req); err != nil {
		RetornarFalha(c, http.StatusBadRequest, "JSON inválido")
		return
	}

	id, err := m.servico.RegistrarUsuario(req)
	if err != nil {
		RetornarFalha(c, http.StatusBadRequest, err.Error())
		return
	}

	RetornarSucesso(c, http.StatusCreated, id)
}

func (m *ManipuladorAutenticacao) ManipularEntrada(c *gin.Context) {
	var req requisicao.RequisicaoUsuario

	if err := c.ShouldBindJSON(&req); err != nil {
		RetornarFalha(c, http.StatusBadRequest, "JSON inválido")
		return
	}

	token, err := m.servico.Entrar(req)
	if err != nil {
		RetornarFalha(c, http.StatusBadRequest, err.Error())
		return
	}

	cfg, err := configuracao.Carregar()
	if err != nil {
		RetornarFalha(c, http.StatusInternalServerError, "Erro ao carregar configuração")
		return
	}

	c.SetCookie(
		"auth",
		token,
		int(cfg.DuracaoJWT.Seconds()),
		"/",
		"",
		true,
		false,
	)

	RetornarSucesso(c, 200, nil)
}
