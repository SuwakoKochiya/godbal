package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xujiajun/godbal"
	"github.com/xujiajun/godbal/driver/mysql"
	"log"
)

func foo() {
	panic("foo func error")
}

func main() {
	database, _ := godbal.NewMysql("root:123@tcp(127.0.0.1:3306)/test?charset=utf8").Open()

	defer database.Close()

	_, err := database.Begin()

	if err != nil {
		log.Fatalln(err)
	}

	defer database.Rollback()

	queryBuilder := mysql.NewQueryBuilder(database)

	rowsAffected, err := queryBuilder.Update("userinfo", "u").Set("u.username", "joe").Set("u.departname", "tecxx").Where("u.uid=?").
		SetParam(4).PrepareAndExecute()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(rowsAffected)

	foo()

	if err := database.Commit(); err != nil {
		log.Fatalln(err)
	}
}
