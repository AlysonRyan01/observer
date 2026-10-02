package modelo

import "time"

type EnderecoLog struct {
	ID string
	EnderecoId string
	Log *string
	DataCriacao time.Time
	CodigoStatus *int
}