package handler

import (
	"fmt"
	"math"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 1000000; i++ {
		math.Pow(36, 89)
	}
	fmt.Fprint(w, "Hello!")
}

func Fibo(w http.ResponseWriter, r *http.Request) {
	l := 1000
	f := make([]int, 0, l)
	for i := 0; i < l; i++ {
		if i <= 1 {
			f = append(f, i + 1)
			continue
		}
		f = append(f, f[i-1] + f[i-2])
	}
	fmt.Fprint(w, "Fibo!")
}
