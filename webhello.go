package main

import (
   "net/http"
   "fmt"
   "strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
  remPartOfURL := r.URL.Path[len("/hello/"):]
  fmt.Fprint(w, "Hello ", remPartOfURL)
}

func shouthelloHandler(w http.ResponseWriter, r *http.Request) {
  remPartOfURL := r.URL.Path[len("/shouthello/"):]
  fmt.Fprint(w, "Hello ", strings.ToUpper(remPartOfURL))
}


func main() {
  http.HandleFunc("/hello/", helloHandler)
  http.HandleFunc("/shouthello/", shouthelloHandler)

  http.ListenAndServe("localhost:9999", nil)
}
