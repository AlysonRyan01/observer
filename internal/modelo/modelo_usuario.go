package modelo

type Usuario struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Senha string `json:"-"`
	Role  Role   `json:"role"`
}
