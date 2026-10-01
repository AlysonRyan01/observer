package roteador

import (
	"go-api/internal/manipulador"

	"github.com/gin-gonic/gin"
)

func AdicionarRotas(r *gin.Engine, m *manipulador.Manipuladores) {
	AdicionarRotasAutenticacao(r, m.Autenticacao)
}
