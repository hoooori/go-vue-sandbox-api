package main

import (
	"go-vue-sandbox-api/models"
	"go-vue-sandbox-api/routers"
)

func main() {
	r := routers.InitRouter()
	db := models.ConnectDB()
	defer db.Close()
	r.Run()
}
