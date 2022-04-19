package main

// 执行存储过程和函数
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

const call = "select * from format_name($1,$2,$3)"

const callMySQL = "CALL simpleproc(?)"

type Result struct {
	Name     string
	Category int
}

func main() {
	db := createConnection()
	defer db.Close()
	r := Result{}

	if err := db.QueryRow(call, "John", "Doe", 32).Scan(&r.Name); err != nil {
		panic(err)
	}
	fmt.Printf("Result is: %+v\n", r)
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
