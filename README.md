```
go clean -testcache && go test ./... -cover
cd cmd
go build -race .
cmd 5 # 5 for example
```
