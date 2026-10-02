package servico

import (
	"context"
	"go-api/internal/configuracao"
	"go-api/internal/modelo"
	"go-api/internal/repositorio"
	"go-api/internal/requisicao"
	"time"

	"github.com/google/uuid"
)

type ServicoEndereco struct {
	repositorio *repositorio.RepositorioEndereco
	configuracao *configuracao.Configuracao
}

func NovoServicoEndereco(r *repositorio.RepositorioEndereco, c *configuracao.Configuracao) *ServicoEndereco {
	return &ServicoEndereco{r, c}
}

func (s *ServicoEndereco) Criar(ctx context.Context, re requisicao.RequisicaoEndereco, usuarioId string, rp requisicao.Paginacao) (string, error) {
	m := &modelo.Endereco{
		ID: uuid.New().String(),
		UsuarioId: usuarioId,
		Endereco: re.Endereco,
		DataCriacao: time.Now().UTC(),
	}

	if err := s.repositorio.Criar(ctx, m); err != nil {
		return "", err
	}

	return m.ID, nil
}

