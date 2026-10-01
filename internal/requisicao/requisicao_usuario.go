package requisicao

import "go-api/internal/modelo"

type RequisicaoUsuario struct {
	Email string      `json:"email" validate:"required,email"`
	Senha string      `json:"senha" validate:"required,min=6"`
	Role  modelo.Role `json:"-"`
}
