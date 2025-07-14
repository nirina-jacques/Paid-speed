package main

import (
	"fmt"
	"net/http"
	"payspeed/cmd/handler"

	_ "github.com/go-sql-driver/mysql"
)

var conf = handler.ServerConfig{}

func main() {
	conf.Config()
	fmt.Println("websocet RUN 9191")
	http.ListenAndServe(":9191", conf.Router)
}
