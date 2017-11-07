package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func main() {
  r, err := http.Get("http://localhost:8080/todos")
  if err != nil { panic(err) }
  defer r.Body.Close()
  body, err := ioutil.ReadAll(r.Body)
  fmt.Println(string(body))
}
