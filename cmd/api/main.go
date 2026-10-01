package main

import (
	"database/sql"
	"go-api/internal/configuracao"
	"go-api/internal/inicializacao"
	"go-api/internal/manipulador"
	"go-api/internal/repositorio"
	"go-api/internal/roteador"
	"go-api/internal/servico"
	"log"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	cfg, err := configuracao.Carregar()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.ObterStringConexao())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	repositorios := repositorio.NovosRepositorios(db)
	servicos := servico.NovosServicos(repositorios, cfg)
	manipuladores := manipulador.NovosManipuladores(servicos)

	if err := inicializacao.InicializarAdmin(cfg, servicos); err != nil {
		log.Fatal(err)
	}

	roteador.AdicionarRotas(app, manipuladores)

	app.Run(":8080")
}
