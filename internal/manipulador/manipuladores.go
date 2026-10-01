package manipulador

import (
	"go-api/internal/servico"
)

type Manipuladores struct {
	Autenticacao *ManipuladorAutenticacao
}

func NovosManipuladores(s *servico.Servicos) *Manipuladores {
	return &Manipuladores{
		Autenticacao: NovoManipuladorAutenticacao(s.Autenticacao),
	}
}
