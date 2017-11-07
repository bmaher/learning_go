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
    Given("Some Buzz").
    UponReceiving("A boring number").
    WithRequest(dsl.Request {
      Method: "GET",
      Path:   "/1",
    }).
    WillRespondWith(dsl.Response {
      Status: 200,
      Body: `Result: "1"`,
    })

  if err := pact.Verify(func() error {
    boring := fmt.Sprintf("http://localhost:%d/1", pact.Server.Port)
    req, err := http.NewRequest("GET", boring, strings.NewReader(""))
    if err != nil { return err }
    if _, err = http.DefaultClient.Do(req); err != nil { return err }
    return nil
  }); err != nil { t.Fatal(err) }

  pact.
    AddInteraction().
    Given("Some Buzz").
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
    fizzing := fmt.Sprintf("http://localhost:%d/3", pact.Server.Port)
    req, err := http.NewRequest("GET", fizzing, strings.NewReader(""))
    if err != nil { return err }
    if _, err = http.DefaultClient.Do(req); err != nil { return err }
    return nil
  }); err != nil { t.Fatal(err) }

  pact.
    AddInteraction().
    Given("Some Buzz").
    UponReceiving("A buzzing number").
    WithRequest(dsl.Request {
      Method: "GET",
      Path:   "/5",
    }).
    WillRespondWith(dsl.Response {
      Status: 200,
      Body: `Result: "Buzz"`,
    })

  if err := pact.Verify(func() error {
    buzzing := fmt.Sprintf("http://localhost:%d/5", pact.Server.Port)
    req, err := http.NewRequest("GET", buzzing, strings.NewReader(""))
    if err != nil { return err }
    if _, err = http.DefaultClient.Do(req); err != nil { return err }
    return nil
  }); err != nil { t.Fatal(err) }

  pact.
    AddInteraction().
    Given("Some Buzz").
    UponReceiving("A fizzing, buzzing number").
    WithRequest(dsl.Request {
      Method: "GET",
      Path:   "/15",
    }).
    WillRespondWith(dsl.Response {
      Status: 200,
      Body: `Result: "FizzBuzz"`,
    })

  if err := pact.Verify(func() error {
    fb := fmt.Sprintf("http://localhost:%d/15", pact.Server.Port)
    req, err := http.NewRequest("GET", fb, strings.NewReader(""))
    if err != nil { return err }
    if _, err = http.DefaultClient.Do(req); err != nil { return err }
    return nil
  }); err != nil { t.Fatal(err) }

  pact.WritePact()
}
