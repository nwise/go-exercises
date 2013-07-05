package main

import (
    _ "github.com/bmizerany/pq"
    "database/sql"
    "fmt"
    "net/http"
    "log"
)

func programHandler(w http.ResponseWriter, r *http.Request) {

  method := r.Method

  switch {
  case method == "GET":
    log.Printf("Method: GET\n")
    findProgram(w, r)
  case method == "POST":
    log.Printf("Method: POST\n")
    createProgram(w, r)
  case method == "PUT":
    log.Printf("Method: PUT\n")
    updateProgram(w, r)
  case method == "DELETE":
    log.Printf("Method: DELETE\n")
    deleteProgram(w, r)
  }

}

func createProgram(w http.ResponseWriter, r *http.Request) {
}

func updateProgram(w http.ResponseWriter, r *http.Request) {
}

func deleteProgram(w http.ResponseWriter, r *http.Request) {
}

func findProgram(w http.ResponseWriter, r *http.Request) {
  var name string

  db, err := sql.Open("postgres", "user=readerwriter dbname=clients password=secret sslmode=disable")

  if( err != nil ) {
    log.Fatal(err)
  }

  programId := r.URL.Path[len("/programs/"):]
  err = db.QueryRow("SELECT name from programs WHERE program_id = $1", programId).Scan(&name)

  switch {
  case err == sql.ErrNoRows:
    log.Printf("No user with that ID.")
  case err != nil:
    log.Fatal(err)
  default:
    log.Printf("Name is %s\n", name)
  }

  fmt.Fprint(w, "Program Name is: ", name)
}

func main() {
  http.HandleFunc("/programs/", programHandler)
  http.ListenAndServe("localhost:9999", nil)
}

