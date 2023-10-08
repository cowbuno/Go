package main

import (
	"SDP/factory"
	"fmt"
)

func main() {

	fct := factory.DatabaseConnectionFactory{}

	mysql := fct.CreateDatabase("MySQL")
	mysql.Connect()

	mysqlRes, _ := mysql.Query("select")
	fmt.Println(mysqlRes)
	fmt.Println()

	postgreSQL := fct.CreateDatabase("PostgreSql")
	postgreSQL.Connect()

	postgreSQLlRes, _ := postgreSQL.Query("select")
	fmt.Println(postgreSQLlRes)
	fmt.Println()

	mongoDB := fct.CreateDatabase("MongoDB")
	mongoDB.Connect()

	mongoDBRes, _ := mongoDB.Query("select")
	fmt.Println(mongoDBRes)
}
