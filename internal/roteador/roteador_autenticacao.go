package roteador

import (
	"go-api/internal/manipulador"

	"github.com/gin-gonic/gin"
)

func AdicionarRotasAutenticacao(r *gin.Engine, m *manipulador.ManipuladorAutenticacao) {
	grupo := r.Group("/auth")

	grupo.POST("/login", m.ManipularEntrada)
	grupo.POST("/register", m.ManipularRegistro)
}
