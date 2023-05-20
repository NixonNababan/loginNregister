package main

import (
	"loginnregister/config"
	"loginnregister/lib/seeder"
	route "loginnregister/routes"
	"loginnregister/util"

	"github.com/go-playground/validator/v10"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	config.InitDB()
	config.InitialMigration()
	seeder.DBSeed(config.DB)
}

func main() {

	e := route.New()
	e.Validator = &util.CustomValidator{
		Validator: validator.New(),
	}
	e.Logger.Fatal(e.Start(":8080"))
}
