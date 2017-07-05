package main

import (
    "database/sql"
    _ "github.com/Go-SQL-Driver/MySQL"
    "fmt"
)

const (
    DB_HOST = "tcp(127.0.0.1:3306)"
    DB_NAME = "ocsdb"
    DB_USER = "ocsuser"
    DB_PASS = "ocspass"
)

func main() {
    dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("Failed to run query", err)
        return
    }
    rows, err := db.Query("select * from accountinfo")
    if err != nil {
        fmt.Println("Failed to run query", err)
        return
    }
    defer rows.Close()
    cols, err := rows.Columns()
    if err != nil {
        fmt.Println("Failed to get columns", err)
        return
    }

    rawResult := make([][]byte, len(cols))
    result := make([][]string,  len(cols))
    for i := range result {
        result[i] = make([]string, 1000) // moust know how many Next cows
    }

    dest := make([]interface{}, len(cols))
    for i, _ := range rawResult {
        dest[i] = &rawResult[i]
    }
    var k int
    k = -1
    for rows.Next() {
        k = k +1
        err = rows.Scan(dest...)
        if err != nil {
            fmt.Println("Failed to scan row", err)
            return
        }

        for i, raw := range rawResult {
            if raw == nil {
                result[i][k] = "\\N"
            } else {
                result[i][k] = string(raw)
            }
        }

        for i:= 0; i < len(cols) ; i++ {
               fmt.Print(result[i][k])
               fmt.Print(" ")
        }
        fmt.Println(" ")
    }
}
