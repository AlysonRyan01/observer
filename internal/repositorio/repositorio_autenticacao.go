package repositorio

import (
	"database/sql"
	"go-api/internal/modelo"
)

type RepositorioAutenticacao struct {
	db *sql.DB
}

func NovoRepositorioAutenticacao(db *sql.DB) *RepositorioAutenticacao {
	return &RepositorioAutenticacao{db: db}
}

func (r *RepositorioAutenticacao) Criar(u modelo.Usuario) error {
	_, err := r.db.Exec("INSERT INTO usuarios (id, email, senha, role) VALUES ($1, $2, $3, $4)", u.ID, u.Email, u.Senha, u.Role)

	return err
}

func (r *RepositorioAutenticacao) BuscarPorEmail(email string) (modelo.Usuario, error) {
	var u modelo.Usuario
	err := r.db.QueryRow(
		"SELECT id, email, senha, role FROM usuarios WHERE email = $1",
		email,
	).Scan(&u.ID, &u.Email, &u.Senha, &u.Role)
	return u, err
}
