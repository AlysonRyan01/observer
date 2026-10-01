package servico

import (
	"errors"
	"fmt"
	"go-api/internal/configuracao"
	"go-api/internal/modelo"
	"go-api/internal/repositorio"
	"go-api/internal/requisicao"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

type ServicoAutenticacao struct {
	repositorio  *repositorio.RepositorioAutenticacao
	configuracao *configuracao.Configuracao
}

func NovoServicoAutenticacao(r *repositorio.RepositorioAutenticacao, c *configuracao.Configuracao) *ServicoAutenticacao {
	return &ServicoAutenticacao{
		repositorio:  r,
		configuracao: c,
	}
}

func (s *ServicoAutenticacao) Entrar(req requisicao.RequisicaoUsuario) (string, error) {
	if err := validate.Struct(req); err != nil {
		return "", err
	}

	usuario, err := s.repositorio.BuscarPorEmail(req.Email)
	if err != nil {
		return "", errors.New("Credenciais inválidas")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(usuario.Senha),
		[]byte(req.Senha),
	)
	if err != nil {
		return "", errors.New("Credenciais inválidas")
	}

	agora := time.Now()

	claims := modelo.ClaimsUsuario{
		Role: usuario.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   usuario.ID,
			IssuedAt:  jwt.NewNumericDate(agora),
			ExpiresAt: jwt.NewNumericDate(agora.Add(s.configuracao.DuracaoJWT)),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString([]byte(s.configuracao.SegredoJWT))
}

func (s *ServicoAutenticacao) RegistrarUsuario(req requisicao.RequisicaoUsuario) (string, error) {
	return s.registrar(req, modelo.RoleUsuario)
}

func (s *ServicoAutenticacao) RegistrarAdministrador(req requisicao.RequisicaoUsuario) (string, error) {
	return s.registrar(req, modelo.RoleAdmin)
}

func (s *ServicoAutenticacao) registrar(req requisicao.RequisicaoUsuario, role modelo.Role) (string, error) {
	if err := validate.Struct(req); err != nil {
		return "", err
	}

	_, err := s.repositorio.BuscarPorEmail(req.Email)
	if err == nil {
		return "", errors.New("Usuário existente")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Senha), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	usuario := modelo.Usuario{
		ID:    uuid.New().String(),
		Email: req.Email,
		Senha: string(hash),
		Role:  role,
	}

	if err := s.repositorio.Criar(usuario); err != nil {
		return "", err
	}

	return usuario.ID, nil
}

func (s *ServicoAutenticacao) VerificarToken(tokenString string) (*modelo.ClaimsUsuario, error) {
	claims := &modelo.ClaimsUsuario{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}
		return s.configuracao.SegredoJWT, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("token inválido")
	}

	return claims, nil
}
