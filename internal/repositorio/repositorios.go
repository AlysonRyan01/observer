package repositorio

import "database/sql"

type Repositorios struct {
	Autenticacao *RepositorioAutenticacao
	Endereco *RepositorioEndereco
}

func NovosRepositorios(db *sql.DB) *Repositorios {
	return &Repositorios{
		Autenticacao: NovoRepositorioAutenticacao(db),
		Endereco: NovoRepositorioEndereco(db),
	}
}
