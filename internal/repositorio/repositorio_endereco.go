package repositorio

import (
	"context"
	"database/sql"
	"go-api/internal/modelo"
	"go-api/internal/requisicao"
)

type RepositorioEndereco struct {
	db *sql.DB
}

func NovoRepositorioEndereco(db *sql.DB) *RepositorioEndereco {
	return &RepositorioEndereco{db:db}
}

func (r *RepositorioEndereco) Criar(ctx context.Context, e *modelo.Endereco) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO enderecos (id, usuario_id, endereco, data_criacao) 
		VALUES ($1, $2, $3, $4)`, 
		e.ID, e.UsuarioId, e.Endereco, e.DataCriacao)

	return err
}

func (r *RepositorioEndereco) Listar(ctx context.Context, p requisicao.Paginacao, usuarioId string) ([]modelo.Endereco, error) {
	offset := (p.Pagina - 1) * p.TamanhoPagina

	rows, err := r.db.QueryContext(ctx, `
		SELECT id, usuario_id, endereco, data_criacao 
		FROM enderecos 
		WHERE usuario_id = $1 
		ORDER BY data_criacao, id
		LIMIT $2 OFFSET $3`, usuarioId, p.TamanhoPagina, offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	enderecos := make([]modelo.Endereco, 0)

	for rows.Next() {
		var endereco modelo.Endereco

		if err := rows.Scan(&endereco.ID, &endereco.UsuarioId, 
			&endereco.Endereco, &endereco.DataCriacao); err != nil {
			return enderecos, err
		}

		enderecos = append(enderecos, endereco)
	}

	if err = rows.Err(); err != nil {
		return enderecos, err
	}

	return enderecos, nil
}

func (r *RepositorioEndereco) CriarLog(ctx context.Context, l *modelo.EnderecoLog) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO endereco_logs (id, endereco_id, log, data_criacao, codigo_status) 
		VALUES ($1, $2, $3, $4, $5)`, 
	l.ID, l.EnderecoId, l.Log, l.DataCriacao, l.CodigoStatus)

	return err
}

func (r *RepositorioEndereco) ListarLogs(ctx context.Context, p requisicao.Paginacao, enderecoId string, usuarioId string) ([]modelo.EnderecoLog, error) {
	offset := (p.Pagina - 1) * p.TamanhoPagina

	logs := make([]modelo.EnderecoLog, 0)

	rows, err := r.db.QueryContext(ctx, `
		SELECT l.id, l.endereco_id, l.log, l.data_criacao, l.codigo_status
		FROM endereco_logs l
		JOIN enderecos e ON l.endereco_id = e.id
		WHERE e.id = $1 AND e.usuario_id = $2
		ORDER BY l.data_criacao DESC, l.id
		LIMIT $3 OFFSET $4`,
		enderecoId, usuarioId, p.TamanhoPagina, offset,
	)
	if err != nil {
		return logs, err
	}
	defer rows.Close()

	for rows.Next() {
		var log modelo.EnderecoLog

		if err = rows.Scan(&log.ID, &log.EnderecoId, &log.Log, &log.DataCriacao, &log.CodigoStatus); err != nil {
			return logs, err
		}
		
		logs = append(logs, log)
	}

	if err = rows.Err(); err != nil {
		return logs, err 
	}

	return logs, nil
}