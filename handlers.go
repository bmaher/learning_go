package main

import (
  "encoding/json"
  "io"
  "io/ioutil"
  "fmt"
  "net/http"

  "github.com/gorilla/mux"
)

const ONE_MB int64 = 1048576
const UNPROCESSABLE_ENTITY int = 422

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json;charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err:= json.NewEncoder(w).Encode(todos); err != nil { panic(err) }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  todoID := vars["todoID"]
  fmt.Fprintf(w, "Todo show: %q", todoID)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
  var todo Todo
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, ONE_MB))
  if err != nil { panic(err) }
  if err := r.Body.Close(); err != nil { panic(err) }
  if err := json.Unmarshal(body, &todo); err != nil {
    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(UNPROCESSABLE_ENTITY)
    if err := json.NewEncoder(w).Encode(err); err != nil { panic(err) }
  }

  t := RepoCreateTodo(todo)
  w.Header().Set("Content-Type", "application/json;charset=UTF-8")
  w.WriteHeader(http.StatusCreated)
  if err := json.NewEncoder(w).Encode(t); err != nil { panic(err) }
}
