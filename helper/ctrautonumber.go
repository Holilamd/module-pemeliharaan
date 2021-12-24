package helper

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

func Ctrautonumber(db *pgx.Conn, table string, header string) int {
	var next_counter int
	sql := "select next_counter from ctrautonumber where table_name=$1 "
	err := db.QueryRow(context.Background(), sql, table).Scan(&next_counter)
	if err != nil {
		log.Printf("QueryRow() error: %v.", err)
	}
	if next_counter == 0 {
		sql2 := "insert into ctrautonumber (table_name,header,next_counter) values($1,$2,$3)"
		db.Exec(context.Background(), sql2, table, header, 1)
		next_counter = 1
		return int(next_counter)

	} else {
		sql2 := "update  ctrautonumber  set next_counter=next_counter+1 where table_name=$1 "
		db.Exec(context.Background(), sql2, table)
		return int(next_counter + 1)
	}

}
