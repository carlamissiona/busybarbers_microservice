package database

import (
	"database/sql"
	_ "encoding/json"
	_ "fmt"
	"log"
	_ "net/http"
	"os"
	_ "os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func OpeDatabase() *sql.DB {

	err := godotenv.Load(".env")


	if err != nil {
		log.Fatalf("Error loading .env file")
	  }
	
	
    urldb := os.Getenv("POSTGRES_URL")
	Database, err := sql.Open("postgres", urldb)
	if err != nil {
		log.Println("Error after OS env file")
		panic(err)
	}
	return Database
}
func SetupDatabase() *sql.DB {
	var err error
    urldb := os.Getenv("POSTGRES_URL")
	if err != nil {
		log.Fatalf("failed reading env file: %v", err)
	}
    
	log.Println("os env file")
	Database, err := sql.Open("postgres", urldb)
	if err != nil {
		log.Println("Error after OS env file")
		panic(err)
	}
	log.Printf("DB connection %v", Database)
	err = Database.Ping();log.Println("Passed Ping & Error after OS env file");
	if err != nil {
		 
		Database, err = sql.Open("postgres", urldb)
		if err != nil {
			log.Println("Error after OS env file")
		   
		}
		log.Printf("failed No DB connection %v", err)
	}

	rows, err := Database.Query("SELECT * FROM bbr_articles")
	 
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("db_instance") 

	log.Println(Database)
	log.Println(rows)
	return Database
}
