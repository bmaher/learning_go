## Install go

https://golang.org/doc/install

## Download this repo

> N.B. Go requires projects to be on the $GOPATH

```bash
git clone git@github.com:bmaher/learning_go.git
```

## Install pact-go

https://github.com/pact-foundation/pact-go

## Run the pact daemon

```bash
pact-go daemon
```

## Install dependencies

```bash
go get github.com/gorilla/mux
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
