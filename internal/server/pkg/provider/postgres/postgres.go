package postgres

import (
	"database/sql"

	"github.com/labstack/gommon/log"
)

type Provider struct {
	db *sql.DB
}

func New(databaseDSN string) *Provider {
	db, err := sql.Open("postgres", databaseDSN)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	if err = createTable(db); err != nil {
		log.Fatal(err)
	}
	return &Provider{db: db}
}

func (p *Provider) Close() error {
	return p.db.Close()
}

func createTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id text primary key,
		login text not null unique,
		password text not null
    );
	CREATE TABLE IF NOT EXISTS auth_data (
		id serial primary key,
		user_id text not null references users(id),
		login text not null,
		password text not null,
		metadata text
	);
	CREATE TABLE IF NOT EXISTS text_data (
		id serial primary key,
		user_id text not null references users(id),
		"data" text not null,
		metadata text
	);
	CREATE TABLE IF NOT EXISTS blob_data (
		id serial primary key,
		user_id text not null references users(id),
		"data" bytea not null,
		metadata text
	);
	CREATE TABLE IF NOT EXISTS card_data (
		id serial primary key,
		user_id text not null references users(id),
		card_number text not null unique,
		"month" text not null,
		"year" text not null,
		cvc text not null,
		"name" text,
		surname text,
		metadata text
	);`
	_, err := db.Exec(query)
	return err
}
