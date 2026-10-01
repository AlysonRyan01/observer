package middleware

import (
	"go-api/internal/manipulador"
	"go-api/internal/servico"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Autenticacao(s *servico.ServicoAutenticacao) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("auth")
		if err != nil {
			manipulador.RetornarFalha(c, http.StatusUnauthorized, "Usuário não autenticado")
			return
		}

		claims, err := s.VerificarToken(cookie.Value)
		if err != nil {
			manipulador.RetornarFalha(c, http.StatusUnauthorized, "Erro na autenticação")
			return
		}

		c.Set("UsuarioId", claims.Subject)
		c.Next()
	}
}
