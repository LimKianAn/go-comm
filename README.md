### Build from the local repository
```
git clone https://github.com/LimKianAn/go-comm.git
cd cmd/comm
go build -race .
comm <tx-number> # for example "comm 5" or "comm.exe 5" (windows)
```

### Run tests
```
go clean -testcache && go test ./... -cover
```

### Build from the remote repository
```
go get github.com/LimKianAn/go-comm/cmd/comm # The generated binary is in $GOPATH/bin
comm <tx-number> # 
```