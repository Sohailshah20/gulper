package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/sohailshah20/csvbatch/csv"
)

type Db struct {
	conn *pgx.Conn
}

func NewDb(dbUri string) (*Db, error) {
	conn, err := pgx.Connect(context.Background(), dbUri)
	if err != nil {
		return nil, err
	}
	return &Db{
		conn: conn,
	}, nil
}

func (db *Db) BatchInsert(columns []string, data [][]string) {
	defer db.conn.Close(context.Background())
	tableName := "customers"
	str, strArr := csv.FormatValues(columns)
	insertStr := fmt.Sprintf(`INSERT INTO %s (%s) Values (%s)`, tableName, csv.GetRowString(columns), str)
	batch := &pgx.Batch{}
	for i := 1; i < len(data); i++ {
		args := csv.BatchArggs(strArr, data[i])
		batch.Queue(insertStr, args)
	}
	results := db.conn.SendBatch(context.Background(), batch)
	defer results.Close()
	for i := 1; i < len(data); i++ {
		_, err := results.Exec()
		if err != nil {
			fmt.Println("err ", err)
		}
	}
}
