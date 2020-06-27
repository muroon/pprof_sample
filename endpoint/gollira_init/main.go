package main

import (
    "github.com/gorilla/mux"
    "github.com/muroon/pprof_sample/handler"
    "net"

    "net/http"
    _ "net/http/pprof"
)

func main() {
    r := mux.NewRouter()
    r.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
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
