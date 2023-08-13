package db

import (
	"database/sql"
	"go-gin-crud-auth/utils"
	"log"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() {
	cfg := mysql.Config{
		User:   "root", // TODO: replace with env
		Passwd: "root", // TODO: replace with env
		Net:    "tcp",
		Addr:   "127.0.0.1:3306", // TODO: replace with env
		DBName: "powerlifterplpus",
	}

	var err error
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		utils.Error.Report(err)
		log.Fatal(err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		utils.Error.Report(pingErr)
		log.Fatal(pingErr)
	}
}

func SelectMultiple[T any](mapLine func(*sql.Rows, *T) error, query string, args ...any) ([]*T, error) {
	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, utils.Error.Report(err)
	}
	defer rows.Close()

	list := []*T{}

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var t T
		if err := mapLine(rows, &t); err != nil {
			return nil, utils.Error.Report(err)
		}
		list = append(list, &t)
	}
	if err := rows.Err(); err != nil {
		return nil, utils.Error.Report(err)
	}
	return list, nil
}
func SelectSingle[T any](mapLine func(*sql.Rows, *T) error, query string, args ...any) (*T, error) {
	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, utils.Error.Report(err)
	}
	defer rows.Close()

	for rows.Next() {
		var t T
		if err := mapLine(rows, &t); err != nil {
			return nil, utils.Error.Report(err)
		}
		return &t, nil
	}
	if err := rows.Err(); err != nil {
		return nil, utils.Error.Report(err)
	}
	return nil, nil
}

func Insert(query string, args ...any) (int, error) {
	result, err := DB.Exec(query, args...)
	if err != nil {
		return 0, utils.Error.Report(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, utils.Error.Report(err)
	}

	return int(id), nil
}

func Update(query string, args ...any) error {
	_, err := DB.Exec(query, args...)
	if err != nil {
		return utils.Error.Report(err)
	}

	return nil
}

func Delete(query string, args ...any) error {
	_, err := DB.Exec(query, args...)
	if err != nil {
		return utils.Error.Report(err)
	}

	return nil
}
