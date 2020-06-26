package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "math"
    "net"

    "net/http"
)
import "net/http/pprof"

func AttachProfiler(router *mux.Router) {
    router.HandleFunc("/debug/pprof/", pprof.Index)
    router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
    router.HandleFunc("/debug/pprof/profile", pprof.Profile)
    router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
}

func SayHello(w http.ResponseWriter, r *http.Request) {
    for i := 0; i < 1000000; i++ {
        math.Pow(36, 89)
    }
    fmt.Fprint(w, "Hello!")
}

func main() {
    r := mux.NewRouter()
    AttachProfiler(r)
    r.HandleFunc("/hello", SayHello)
    lin, err := net.Listen("tcp4", ":8080")
    if err != nil {
        panic(err)
    }
    defer lin.Close()
    s := new(http.Server)
    s.Handler = r
    s.Serve(lin) 
}
