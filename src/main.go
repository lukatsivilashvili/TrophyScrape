package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	server := NewApiServer(":3000")
	server.Run()

}
