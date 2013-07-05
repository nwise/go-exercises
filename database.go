package main

import (
    _ "github.com/bmizerany/pq"
    "database/sql"
    "fmt"
)


func main() {
  db, err := sql.Open("postgres", "user=readerwriter dbname=clients password=secret sslmode=disable")

  fmt.Println("Connection: ", db)
  fmt.Println("Errors: ", err)

  rows, err := db.Query("SELECT name from clients")

  for rows.Next() {
    var name string
    if err := rows.Scan(&name); err != nil {
      fmt.Println("Error")
    }
    fmt.Printf("Client: %s\n", name)
  }

}
