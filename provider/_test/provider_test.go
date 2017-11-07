package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

var dir, _ = os.Getwd()
var pactDir = fmt.Sprintf("%s/../../consumer/_test/pacts", dir)
var pactFile = fmt.Sprintf("%s/fizzer-buzzer.json", pactDir)
var logDir = fmt.Sprintf("%s/log", dir)
var port = 8080

func TestPact_Provider(t *testing.T) {
	pact := createPact()

	err := pact.VerifyProvider(types.VerifyRequest{
		ProviderBaseURL: fmt.Sprintf("http://localhost:%d", port),
		PactURLs:        []string{filepath.ToSlash(pactFile)},
	})

	if err != nil {
		t.Fatal("Error:", err)
	}
}

func createPact() dsl.Pact {
	return dsl.Pact{
		Port:     6666,
		Consumer: "Fizzer",
		Provider: "Buzzer",
		LogDir:   logDir,
		PactDir:  pactDir,
	}
}
