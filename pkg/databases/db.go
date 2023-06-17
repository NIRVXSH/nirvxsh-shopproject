package databases

import (
	"log"

	"github.com/NIRVXSH/NIRVXSH-shop-project/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func DbConnect(cfg config.IDbConfig) *sqlx.DB {
	db, err := sqlx.Connect("pgx", cfg.Url())
	if err != nil {
		log.Fatalf("connect to db failed %v\n", err)
	}
	db.DB.SetMaxOpenConns(cfg.MaxOpenConns())
	return db
}
