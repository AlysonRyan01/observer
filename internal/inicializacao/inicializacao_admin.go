package inicializacao

import (
	"errors"
	"go-api/internal/configuracao"
	"go-api/internal/modelo"
	"go-api/internal/requisicao"
	"go-api/internal/servico"
)

func InicializarAdmin(c *configuracao.Configuracao, s *servico.Servicos) error {
	if c.SenhaAdmin == "" || c.EmailAdmin == "" {
		return errors.New("Email ou senha do admin não configurados")
	}

	servicoAutenticacao := s.Autenticacao

	req := requisicao.RequisicaoUsuario{
		Email: c.EmailAdmin,
		Senha: c.SenhaAdmin,
		Role:  modelo.RoleAdmin,
	}

	_, err := servicoAutenticacao.RegistrarAdministrador(req)
	if err == nil {
		return nil
	}

	if err.Error() == "Usuário existente" {
		return nil
	}

	return err
}
