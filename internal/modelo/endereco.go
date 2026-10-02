package modelo

import "time"

type Endereco struct {
	ID string
	UsuarioId string
	Endereco string
	Logs []EnderecoLog
	DataCriacao time.Time
}