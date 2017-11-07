package main

import (
  "fmt"
  "net/http"

  "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Get ready to be buzzed!")
}

func FizzBuzz(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  num := vars["number"]
  res := num

  var i int
  if _, err := fmt.Sscan(num, &i); err == nil {
    if i % 15 == 0 {
      res = "FizzBuzz"
    } else if i % 3 == 0 {
      res = "Fizz"
    } else if i % 5 == 0 {
      res = "Buzz"
    }
    fmt.Fprintf(w, "Result: %q", res)
  }
}
