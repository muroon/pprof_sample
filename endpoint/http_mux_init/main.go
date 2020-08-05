package main

import (
    "github.com/muroon/pprof_sample/handler"
    "net"

    "net/http"
    _ "net/http/pprof"
)

func AttachProfiler(router *http.ServeMux) {
    router.Handle("/debug/pprof/", http.DefaultServeMux)
}

func main() {
    r := http.NewServeMux() // Custom Mux
    AttachProfiler(r)
    r.HandleFunc("/", handler.Hello)
    r.HandleFunc("/hello", handler.Hello)
    r.HandleFunc("/fibo", handler.Fibo)
    lin, err := net.Listen("tcp4", ":8080")
    if err != nil {
        panic(err)
    }
    defer lin.Close()
    s := new(http.Server)
    s.Handler = r
    s.Serve(lin) 
}
