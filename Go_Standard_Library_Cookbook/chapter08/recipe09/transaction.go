package main

// 处理事务
import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const selOne = "SELECT id,title,content FROM post WHERE ID = $1;"
const insert = "INSERT INTO post(ID,TITLE,CONTENT) VALUES (4,'Transaction Title','Transaction Content');"

type Post struct {
	ID      int
	Title   string
	Content string
}

func main() {
	db := createConnection()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	_, err = tx.Exec(insert)
	if err != nil {
		panic(err)
	}
	p := Post{}
	// 在其它会话/事务中查询
	if err := tx.QueryRow(selOne, 4).Scan(&p.ID, &p.Title, &p.Content); err != nil {
		fmt.Println("Got error for db.Query:" + err.Error())
	}
	fmt.Println(p)
	// 事务内查询
	if err := tx.QueryRow(selOne, 4).Scan(&p.ID, &p.Title, &p.Content); err != nil {
		fmt.Println("Got error for db.Query:" + err.Error())
	}
	fmt.Println(p)
	// 提交或回滚之后
	// 事务需要重新创建
	err = tx.Rollback()
	if err != nil {
		panic(err)
	}

	fmt.Println("\nTransaction with context")
	ctx, canc := context.WithCancel(context.Background())
	tx, err = db.BeginTx(ctx, &sql.TxOptions{sql.LevelReadUncommitted, false})
	if err != nil {
		panic(err)
	}
	_, err = tx.Exec(insert)
	if err != nil {
		panic(err)
	}
	canc()
	err = tx.Commit()
	if err != nil {
		fmt.Println("Error occured:" + err.Error())
	}
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
