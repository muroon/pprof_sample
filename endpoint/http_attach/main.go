package main

import (
    "github.com/muroon/pprof_sample/handler"
    "net"

    "net/http"
    "net/http/pprof"
)

func AttachProfiler() {
    http.HandleFunc("/debug/pprof/", pprof.Index)
    http.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
    http.HandleFunc("/debug/pprof/profile", pprof.Profile)
    http.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
}

func main() {
    AttachProfiler()
    http.HandleFunc("/", handler.Hello)
    http.HandleFunc("/hello", handler.Hello)
    http.HandleFunc("/fibo", handler.Fibo)
    lin, err := net.Listen("tcp4", ":8080")
    if err != nil {
        panic(err)
    }
    defer lin.Close()
    s := new(http.Server)
    s.Handler = http.DefaultServeMux
    s.Serve(lin) 
}
