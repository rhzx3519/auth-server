package mysql

import (
    "database/sql"
    "fmt"
    "github.com/go-sql-driver/mysql"
    "log"
    "os"
)

var (
    DB *sql.DB
)

func getenv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func InitDB() {
    // InitDB config
    cfg := mysql.Config{
        User:                 getenv("DBUSER", "root"),
        Passwd:               getenv("DBPASS", ""),
        Net:                  "tcp",
        Addr:                 fmt.Sprintf("%s:%s", getenv("DBHOST", "127.0.0.1"), getenv("DBPORT", "3306")),
        DBName:               "demo-brokers",
        AllowNativePasswords: true,
        ParseTime:            true,
    }
    fmt.Println(cfg.FormatDSN())
    // Get a DB handler
    var err error
    DB, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }
    pingErr := DB.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Mysql Connected...")
}

func PostDB() {
    if err := DB.Close(); err != nil {
        log.Panicln(err)
    }
}
