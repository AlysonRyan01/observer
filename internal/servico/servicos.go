package servico

import (
	"go-api/internal/configuracao"
	"go-api/internal/repositorio"
)

type Servicos struct {
	Autenticacao *ServicoAutenticacao
}

func NovosServicos(r *repositorio.Repositorios, c *configuracao.Configuracao) *Servicos {
	return &Servicos{
		Autenticacao: NovoServicoAutenticacao(r.Autenticacao, c),
	}
}
