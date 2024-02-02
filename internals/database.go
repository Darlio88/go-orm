package internals

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func DatabaseInstance() *sql.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("fatal error reading .env file", err)

	}
	dsn := os.Getenv("DB_URL")
	// log.Println(dsn)
	if dsn == "" {
		log.Println("the database url has not been set")

	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY,name VARCHAR(255),age INT)")
	if err != nil {
		log.Println("failed to seed database")
	}

	return db
}
