package main

import (
	"github.com/AnushaSankaranarayanan/inventoryservice/receipt"
	"github.com/AnushaSankaranarayanan/inventoryservice/database"
	"github.com/AnushaSankaranarayanan/inventoryservice/product"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

const apiBasePath = "/api"

func main() {
	database.SetupDatabase()
	receipt.SetupRoutes(apiBasePath)
	product.SetupRoutes(apiBasePath)
	http.ListenAndServe(":5000", nil)
}
