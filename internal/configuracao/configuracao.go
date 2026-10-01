package configuracao

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Configuracao struct {
	Porta      string
	HostBD     string
	PortaBD    string
	UsuarioBD  string
	SenhaBD    string
	NomeBD     string
	SegredoJWT string
	DuracaoJWT time.Duration
	EmailAdmin string
	SenhaAdmin string
}

func Carregar() (*Configuracao, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("erro ao carregar .env: %w", err)
	}

	config := &Configuracao{
		Porta:      obterVariavelAmbiente("PORTA", "8080"),
		HostBD:     obterVariavelAmbiente("HOST_BD", "localhost"),
		PortaBD:    obterVariavelAmbiente("PORTA_BD", "5432"),
		UsuarioBD:  obterVariavelAmbiente("USUARIO_BD", ""),
		SenhaBD:    obterVariavelAmbiente("SENHA_BD", ""),
		NomeBD:     obterVariavelAmbiente("NOME_BD", ""),
		SegredoJWT: obterVariavelAmbiente("SEGREDO_JWT", ""),
		DuracaoJWT: 24 * time.Hour,
		EmailAdmin: obterVariavelAmbiente("EMAIL_ADMIN", ""),
		SenhaAdmin: obterVariavelAmbiente("SENHA_ADMIN", ""),
	}

	if config.UsuarioBD == "" || config.SenhaBD == "" {
		return nil, fmt.Errorf("UsuarioBD ou SenhaBD não configurados")
	}

	return config, nil
}

func (c *Configuracao) ObterStringConexao() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.HostBD,
		c.PortaBD,
		c.UsuarioBD,
		c.SenhaBD,
		c.NomeBD,
	)
}

func obterVariavelAmbiente(chave, padrao string) string {
	if valor, ok := os.LookupEnv(chave); ok {
		return valor
	}

	return padrao
}
