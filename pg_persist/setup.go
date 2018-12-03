package pg_persist

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

var Db *sql.DB

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func ConnectToPGDB(cfg Config) (err error) {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.User, cfg.Password, cfg.Database, cfg.Host, cfg.Port))
	if err != nil {
		err = errors.Wrapf(err, "Couldn't open connection to postgre database (%s)", cfg.Database)
		return
	}
	if err = db.Ping(); err != nil {
		err = errors.Wrapf(err, "Couldn't ping postgre database (%s)", cfg.Database)
		return
	}
	Db = db
	return
}
