package main

import (
	"fmt"
	"net/http"
)


type Logger struct{}


func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method,r.URL.Path)

	http.DefaultServeMux.ServeHTTP(w,r)
}


