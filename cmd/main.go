package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"users/delivery/rest"
	"users/service"

	account "users/domain/account/repository"
	account_mysql "users/domain/account/repository/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	var cfg config

	// load config from .env
	err := LoadAndParse("LEMONILO", &cfg)
	if err != nil {
		log.Println(err)
	}
	log.Println("Config successfully loaded")

	// connect to database
	source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)
	db, err := sqlx.Connect("mysql", source)
	if err != nil {
		log.Println(err)
	}
	log.Println("Database successfully initialized")

	// set max connection
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(60 * time.Second)
	db.SetMaxIdleConns(10)

	// init service
	router := httprouter.New()
	var account account.Repository = account_mysql.New(db)
	service := service.New(account)
	rest.New(service).Routing(router)
	log.Println("Service successfully initialized")

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}

	db.Close()
	log.Println("Database successfully stopped")

}

// LoadAndParse to load and parse config from .env into struct
func LoadAndParse(prefix string, out interface{}, filenames ...string) error {
	err := godotenv.Load(filenames...)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = envconfig.Process(prefix, out)
	if err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
