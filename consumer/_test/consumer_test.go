package main

import (
  "fmt"
  "net/http"
  "testing"
  "strings"

  "github.com/pact-foundation/pact-go/dsl"
)

func TestFizzBuzz(t *testing.T) {
  pact := dsl.Pact {
    Port:     6666,
    Consumer: "Fizzer",
    Provider: "Buzzer",
  }

  defer pact.Teardown()

  pact.
    AddInteraction().
    Given("Some Fizz").
    UponReceiving("A fizzing number").
    WithRequest(dsl.Request {
      Method: "GET",
      Path:   "/3",
    }).
    WillRespondWith(dsl.Response {
      Status: 200,
      Body: `Result: "Fizz"`,
    })

  if err := pact.Verify(func() error {
    u := fmt.Sprintf("http://localhost:%d/3", pact.Server.Port)
    req, err := http.NewRequest("GET", u, strings.NewReader(""))
    if err != nil { return err }
    if _, err = http.DefaultClient.Do(req); err != nil { return err }
    return nil
  }); err != nil { t.Fatal(err) }
  pact.WritePact()
}
