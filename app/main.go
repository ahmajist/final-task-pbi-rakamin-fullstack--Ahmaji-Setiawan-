package main

import (
	"log"
	"net/http"
	"FINALTASK-BTPN-SYARIAH/database"
	"FINALTASK-BTPN-SYARIAH/routers"
)

func main() {
	database.InitDB()

	//endpoint API
	routers.DefineRoutes()

	log.Fatal(http.ListenAndServe(":3306", nil))
}
