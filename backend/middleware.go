package main

import (
	"fmt"
	"net/http"
)


type Logger {}


func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(r.Method,r.URL.Path)

	http.DefaultServeMux.ServeHTTP(w,r)
}


