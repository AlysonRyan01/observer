package manipulador

import "github.com/gin-gonic/gin"

type Resposta struct {
	Dados   any    `json:"data"`
	Erro    string `json:"erro"`
	Sucesso bool   `json:"isSuccess"`
}

func RetornarSucesso(c *gin.Context, codigo int, dados any) {
	resposta := Resposta{
		Dados:   dados,
		Erro:    "",
		Sucesso: true,
	}

	c.JSON(codigo, resposta)
}

func RetornarFalha(c *gin.Context, codigo int, erro string) {
	resposta := Resposta{
		Dados:   nil,
		Erro:    erro,
		Sucesso: false,
	}

	c.JSON(codigo, resposta)
}
