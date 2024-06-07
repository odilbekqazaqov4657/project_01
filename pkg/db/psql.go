package db

import (
	config "app/cofig"
	"context"
	"fmt"

	"github.com/jackc/pgx"
)

var ctx = context.Background()

func Conn(cfg config.Config) (*pgx.Conn, error) {

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.DbUser,
		cfg.DbPassword,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName,
	)

	return pgx.Connect(ctx, dbURL)

}
