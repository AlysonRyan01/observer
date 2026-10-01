package repositorio

import "database/sql"

type Repositorios struct {
	Autenticacao *RepositorioAutenticacao
}

func NovosRepositorios(db *sql.DB) *Repositorios {
	return &Repositorios{
		Autenticacao: NovoRepositorioAutenticacao(db),
	}
}
