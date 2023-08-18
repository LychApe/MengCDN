package main

import (
	"MengCDN/internal/model"
	"MengCDN/internal/router"
)

func main() {
	model.InitDB()
	router.InitRouter()

}
