### Test, build and run from the local repository
```
git clone https://github.com/LimKianAn/go-comm.git
cd go-comm
go clean -testcache && go test ./... -cover
cd cmd/comm
go build -race .
comm <tx-number> # for example "comm 5" or "comm.exe 5" (windows)
```

### Build from the remote repository
```
go get github.com/LimKianAn/go-comm/cmd/comm # The generated binary is in "$GOPATH/bin".
comm <tx-number>
```
