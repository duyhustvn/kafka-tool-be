package sqlclient

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func NewSqlite() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "kafka_tool.db")
	if err != nil {
		return nil, fmt.Errorf("cannot connect to sqlite %+v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("cannot ping sqlite %+v", err)
	}

	var schema = `
	CREATE TABLE IF NOT EXISTS requests (
		id integer primary key,
		title varchar(255) not null,
		topic varchar(255) not null,
		quantity integer default 1,
		type varchar(30) not null,
		message text not null
	);
	`
	if _, err := db.Exec(schema); err != nil {
		return nil, fmt.Errorf("init table failed %+v", err)
	}

	return db, nil
}
