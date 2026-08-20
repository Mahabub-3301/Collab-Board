package main

import (
	"net/http"
	"log"
	"time"
	"crypto/rand"
	"encoding/hex"
)

func generateID() string {
	bytes := make([]byte,16)

	_, err := rand.Read(bytes)
	if err != nil {
		return "unknown"
	}
	return hex.EncodeToString(bytes)
}

func RequestIDMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		requestID := w.Header().Get("X-Request-ID")

		if requestID == "" {
			requestID = generateID()
		}

		ctx := setRequestID(r.Context(),requestID)

		r = r.WithContext(ctx)

		w.Header().Set("X-Request-ID",requestID)
		next.ServeHTTP(w,r)
	})
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		start := time.Now()

		next.ServeHTTP(w,r)

		log.Printf("%s %s %v",r.Method,r.URL.Path,time.Since(start))
	})
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic recovered: %v",err)

				writeError(
					w,
					http.StatusInternalServerError,
					"Internal Server Error",
				)
			}
		}()

		next.ServeHTTP(w,r)
	})
}


func Chain(
	handler http.Handler,
	middleware ...func(http.Handler) http.Handler,
) http.Handler {

	for i := len(middleware)-1; i >= 0; i-- {
		handler = middleware[i](handler)
	}
	return handler
}



