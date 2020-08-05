# pprof_sample

## command

### start listen server

```
go run endpoint/http_init/main.go
```

You can choose [another endpoint](https://github.com/muroon/pprof_sample/tree/master/endpoint) without http_init.

### send requests

```
for i in `seq 1 1000`; do curl http://127.0.0.1:8080/hello; done
```

### pprof

```
go tool pprof -http=":22222" http://127.0.0.1:8080/debug/pprof/profile
```

## reference
https://stackoverflow.com/questions/19591065/profiling-go-web-application-built-with-gorillas-mux-with-net-http-pprof

