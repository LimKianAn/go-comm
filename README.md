```
go clean -testcache && go test ./... -cover
cd cmd/comm
go build -race .
cmd 5 # 5 for example
```
