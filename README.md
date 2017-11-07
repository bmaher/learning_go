## Run

```bash
cd consumer/_test
go test

cd ../../provider
go run ./*.go &
# Allow on Firewall if asked

cd _test
go test
```
