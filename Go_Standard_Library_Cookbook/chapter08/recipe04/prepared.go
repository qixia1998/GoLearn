package main

// 带有准备好的语句的操作
import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const trunc = "TRUNCATE TABLE post;"
const ins = "INSERT INTO post(ID, TITLE, CONTENT) VALUES ($1, $2, $3)"

var testTable = []struct {
	ID      int
	Title   string
	Content string
}{
	{1, "Title One", "Content of title one"},
	{2, "Title Two", "Content of title two"},
	{3, "Title Three", "Content of title three"},
}

func main() {
	db := createConnection()
	defer db.Close()

	// 截断表
	_, err := db.Exec(trunc)
	if err != nil {
		panic(err)
	}

	stm, err := db.Prepare(ins)
	if err != nil {
		panic(err)
	}

	inserted := int64(0)
	for _, val := range testTable {
		fmt.Printf("Inserting record ID: %d\n", val.ID)
		// 执行准备好的声明
		r, err := stm.Exec(val.ID, val.Title, val.Content)
		if err != nil {
			fmt.Printf("Connect insert record ID : %d\n", val.ID)
		}
		if affected, err := r.RowsAffected(); err == nil {
			inserted = inserted + affected
		}
	}

	fmt.Printf("Result: Inserted %d rows.\n", inserted)
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
