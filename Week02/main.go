package main

import (
	"fmt"
	"database/sql"

	"github.com/pkg/errors"
)

const (
	sqlStatement = `SELECT name, email FROM users WHERE id=$1;`

	errNoRows = errors.New("no rows")
)

type dao struct {
	db *sql.DB
}

func main() {
	dao, err := newDAO("test1")
	if err != nil {
		log.Fatal(err)
	}

	id := 666
	err := service(dao, id)
	if err != nil {
		switch {
		case errors.Is(err, errNoRows):
			fmt.Println("no found by id of %v", id)
		default:
			fmt.Println("encounter error: %v", err)
		}
	}
}

func newDAO(name string) (*dao, error) {
	// Opening a driver typically will not attempt to connect to the database.
	db, err := sql.Open("driver-name", fmt.Sprintln("database=%v", name)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		return nil, err
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	return &dao{db: db}, nil
}

func service(d *dao, id int) error {
	name, email, err := d.findByID(id)
	if err != nil {
		return err
	}

	return sendResetMail(name, email)
}

func (d *dao) findByID(id int) (name, email string, e error) {
	row := d.db.QueryRow(sqlStatement, id)
	switch err := row.Scan(&name, &email); err {
		case sql.ErrNoRows:
			return name, email, fmt.Errorf("%w: no found by id %v", errNoRows, id)
		case nil:
			return name, email, nil
		default:
			return name, email, fmt.Errorf("%w: can't find by id", e)
	}
}

func sendResetMail(name, email string) error {
	return errors.New("not implemented sendResetMail func")
}