package main

import (
    "database/sql"
    _ "github.com/Go-SQL-Driver/MySQL"
    "fmt"
    "os"
)

// DB info
const (
    DB_HOST = "tcp(127.0.0.1:3306)"
    DB_NAME = "ocsdb"
    DB_USER = "ocsuser"
    DB_PASS = "ocspass"
)

// for get table count
func checkCount(rows *sql.Rows) (count int) {
        for rows.Next() {
        err:= rows.Scan(&count)
        checkErr("select count fail",err )
    }
    return count
}

// only for err print & exit
func checkErr(w_msg string, err error) {
    if err != nil {
       fmt.Println(w_msg,err)
     os.Exit(2)
     return
    }
}

func main() {
    var len_count int
    dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
    db, err := sql.Open("mysql", dsn)
    checkErr("Failed to run query",err)
    defer db.Close()

    rows_count, err := db.Query("select COUNT(*) as count from accountinfo")
    checkErr("Failed to run query for count", err)
    len_count = checkCount(rows_count)
    defer rows_count.Close()

    rows, err := db.Query("select * from accountinfo")
    checkErr("Failed to run query", err)
    defer rows.Close()

    cols, err := rows.Columns()
    checkErr("Failed to get columns", err)

    rawResult := make([][]byte, len(cols))
    result := make([][]string,  len(cols))

     for i := range result {
        result[i] = make([]string, len_count)
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
        checkErr("Failed to scan row", err)

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
