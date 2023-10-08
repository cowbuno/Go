package main

import (
	"SDP/factory"
	"fmt"
)

func main() {

	mysql := factory.CreateDatabase("MySQL")
	mysql.Connect()

	mysqlRes, _ := mysql.Query("select")
	fmt.Println(mysqlRes)
}
