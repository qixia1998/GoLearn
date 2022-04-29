package main

// 取消待处理查询
import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

const sel = "SELECT * FROM post p CROSS JOIN (SELECT 1 FROM generate_series(1, 10000000)) tbl"

func main() {
	db := createConnection()
	defer db.Close()

	ctx, canc := context.WithTimeout(context.Background(), 20*time.Millisecond)
	rows, err := db.QueryContext(ctx, sel)
	canc() // 取消查询
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		if rows.Err() != nil {
			fmt.Println(rows.Err())
			continue
		}
		count++
	}

	fmt.Printf("%d rows returned\n", count)
}

func createConnection() *sql.DB {
	connStr := "host=localhost port=5432 dbname=posttest user=postgres password=12345678 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
