package main

import (
    "database/sql"
    _ "github.com/Go-SQL-Driver/MySQL"
    "fmt"
)

const (
    DB_HOST = "tcp(127.0.0.1:3306)"
    DB_NAME = "ocsweb"
    DB_USER = "ocs"
    DB_PASS = "ocssecret"
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
    result := make([]string, len(cols))

    dest := make([]interface{}, len(cols)) // A temporary interface{} slice
    for i, _ := range rawResult {
        dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
    }
    for rows.Next() {
        err = rows.Scan(dest...)
        if err != nil {
            fmt.Println("Failed to scan row", err)
            return
        }

        for i, raw := range rawResult {
            if raw == nil {
                result[i] = "\\N"
            } else {
                result[i] = string(raw)
            }
        }

        fmt.Printf("%#v\n", result)
    }
}
