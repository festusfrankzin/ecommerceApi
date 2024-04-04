package db

import (
    "database/sql"
    "fmt"
    _"github.com/go-sql-driver/mysql"
	"github.com/festusfrankzin/ecommerceApi/configs"
)

var GetEnv = config.GetEnv


// ConnectDB function establishes a connection to the MySQL database
func ConnectDB() (*sql.DB, error) {
    username := GetEnv("DB_USER", "root")
    password := GetEnv("DB_PASSWORD", "1407")
    host := GetEnv("DB_HOST", "localhost")
    port := GetEnv("DB_PORT", "3306")
    dbname := GetEnv("DB_NAME", "EcommerceApi")

    dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)
    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        return nil, err
    }

    // Check if the connection is successful
    err = db.Ping()
    if err != nil {
        return nil, err
    }

    fmt.Println("Connected to MySQL database!")
    return db, nil
}
