package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
)

var dir, _ = os.Getwd()
var pactDir = fmt.Sprintf("%s/../../consumer/_test/pacts", dir)
var logDir = fmt.Sprintf("%s/log", dir)
var port, _ = utils.GetFreePort()

func TestPact_Provider(t *testing.T) {
	pact := createPact()

	err := pact.VerifyProvider(types.VerifyRequest{
		ProviderBaseURL:        fmt.Sprintf("http://localhost:%d", 8080),
		PactURLs:               []string{filepath.ToSlash(fmt.Sprintf("%s/fizzer-buzzer.json", pactDir))},
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
