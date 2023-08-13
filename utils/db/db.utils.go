package db

import (
	"database/sql"
	"go-gin-crud-auth/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() {
	cfg := mysql.Config{
		User:   utils.Config.Database.Username,
		Passwd: utils.Config.Database.Password,
		Net:    "tcp",
		Addr:   utils.Config.Database.Host,
		DBName: utils.Config.Database.DatabaseName,
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

const TX_HANDLE_KEY = "TRANSACTION_HANDLE"

func SetTransaction(c *gin.Context, txHandle *sql.Tx) {
	c.Set(TX_HANDLE_KEY, txHandle)
}

func GetTx(c *gin.Context) *sql.Tx {
	if tx, exists := c.Get(TX_HANDLE_KEY); exists {
		return tx.(*sql.Tx)
	}
	return nil
}

func SelectMultiple[T any](txHandle *sql.Tx, mapLine func(*sql.Rows, *T) error, query string, args ...any) ([]*T, error) {
	rows, err := txHandle.Query(query, args...)
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
func SelectSingle[T any](txHandle *sql.Tx, mapLine func(*sql.Rows, *T) error, query string, args ...any) (*T, error) {
	rows, err := txHandle.Query(query, args...)
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

func Insert(txHandle *sql.Tx, query string, args ...any) (int, error) {
	result, err := txHandle.Exec(query, args...)
	if err != nil {
		return 0, utils.Error.Report(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, utils.Error.Report(err)
	}

	return int(id), nil
}

func Update(txHandle *sql.Tx, query string, args ...any) error {
	_, err := txHandle.Exec(query, args...)
	if err != nil {
		return utils.Error.Report(err)
	}

	return nil
}

func Delete(txHandle *sql.Tx, query string, args ...any) error {
	_, err := txHandle.Exec(query, args...)
	if err != nil {
		return utils.Error.Report(err)
	}

	return nil
}
