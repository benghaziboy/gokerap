package db

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "log"
    "os"
)

var (
    Conn *sql.DB
)

func init() {
    fmt.Println("database stuff")
    dburl := os.Getenv("DATABASE_URL")
    fmt.Println(dburl)
    if dburl == "" {
        dbip := os.Getenv("VAGRANT_DB_1_PORT_5432_TCP_ADDR")
        dburl = fmt.Sprintf("user=postgres sslmode=disable host=%s", dbip)
        fmt.Println(dburl)
    }

    var e error
    Conn, e = sql.Open("postgres", dburl)
    if e != nil {
        log.Fatal(e)
    }
}
