## Install pact-go

https://github.com/pact-foundation/pact-go

## Run the pact daemon

```bash
pact-go daemon
```

## Run the pact tests

```bash
cd consumer/_test
go test

cd ../../provider
go run ./*.go &
# Allow on Firewall if asked

cd _test
go test
```
