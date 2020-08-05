package main

import (
    "github.com/muroon/pprof_sample/handler"
    "net"

    "net/http"
    _ "net/http/pprof"
)


func main() {
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
