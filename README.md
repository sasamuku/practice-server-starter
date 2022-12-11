A server that supports Server::Starter

```
go get github.com/lestrrat/go-server-starter/cmd/start_server
```

```
go build main.go
start_server --port 8080 --pid-file app.pid -- ./main
```

```
kill -HUP `cat app.pid`
```
