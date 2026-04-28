package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	conn *pgx.Conn
	ctx  context.Context // special context for database(md waitCancel)
}

func newConn(ctx context.Context, path string) (*pgx.Conn, error) {
	return pgx.Connect(ctx, path)
}

func NewDataBase(ctx context.Context, path string) (*Database, error) {
	conn, err := newConn(ctx, path)

	if err != nil {
		return &Database{}, err
	}

	return &Database{ctx: ctx, conn: conn}, nil
}
