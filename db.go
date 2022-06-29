package main
import (
	"os"
    "fmt"
	"database/sql"
    _ "github.com/lib/pq"
)
// const (
//     DB_USER     = "postgres"
//     DB_PASSWORD = "Ping_Pong"
//     DB_NAME     = "postgres"
// )

// DB set up
func setupDB() *sql.DB {
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBNAME"))
    fmt.Println(dbinfo)

    db, err := sql.Open("postgres", dbinfo)
    //db, err := sql.Open("postgres", cfg.FormatDSN())
    checkErr(err)

    return db
}