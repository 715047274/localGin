package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

var DB *sql.DB

func init() {
	db, err := sql.Open("sqlite3", "./sqliteDemo.db")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("DB ..... Created")
	DB = db
}

func setupRoute() *gin.Engine {
	r := gin.Default()
	r.GET("/getAccounts", getAccounts)
	return r
}

func getAccounts(c *gin.Context) {
	var sql = `SELECT * FROM accounts`
	rows, _ := DB.Query(sql)
	defer rows.Close()

	accounts := make([]Account, 0)
	for rows.Next() {
		singleAccount := Account{}
		_ = rows.Scan(&singleAccount.Id, &singleAccount.Balance, &singleAccount.Currency, &singleAccount.Owner)

		accounts = append(accounts, singleAccount)
	}
	_ = rows.Err()

	fmt.Println(rows)

	c.JSON(http.StatusOK, accounts)
}

func main() {
	serv := setupRoute()
	serv.Run(":8080")
	fmt.Println("hello world")
}

type Account struct {
	Id       int    `db:"id" json:"id"`
	Owner    string `db:"owner" json:"owner"`
	Currency string `db:"currency" json:"currency"`
	Balance  int16  `db:"balance" json:"balance"`
}
