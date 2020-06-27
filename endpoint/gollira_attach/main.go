package main

import (
    "github.com/gorilla/mux"
    "github.com/muroon/pprof_sample/handler"
    "net"

    "net/http"
    "net/http/pprof"
)

func AttachProfiler(router *mux.Router) {
    router.HandleFunc("/debug/pprof/", pprof.Index)
    router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
    router.HandleFunc("/debug/pprof/profile", pprof.Profile)
    router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
    router.HandleFunc("/debug/pprof/heap", pprof.Handler("heap").ServeHTTP)
}

func main() {
    r := mux.NewRouter()
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
