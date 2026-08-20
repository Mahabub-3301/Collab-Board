package main

import "context"

type contextKey string

const requestIDKey contextKey = "requestID"


func setRequestID(ctx context.Context,id string) context.Context {
	return context.WithValue(ctx,requestIDKey,id)
}

func getRequestID(ctx context.Context) string {
	id, ok := ctx.Value(requestIDKey).(string)

	if !ok {
		return ""
	}
	return id 
}
