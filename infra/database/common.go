package database

import (
	"context"
	"database/sql"
	"os"
)
func GetTestConnection() (*sql.DB) {
    dataSource := os.ExpandEnv("${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_DATABASE}?parseTime=true")
    db, err := sql.Open("mysql", dataSource)
    if err!= nil {
        panic(err.Error())
    }
    if err = db.Ping(); err != nil {
		panic(err)
	}
    return db

}

func GetTestTransaction(db *sql.DB, ctx context.Context) *sql.Tx {
    tx, _ := db.BeginTx(ctx, nil)
    return tx
}