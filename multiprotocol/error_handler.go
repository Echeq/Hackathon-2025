package multiprotocol

import (
    "context"
    "net/http"
)

func ErrorHandlerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusInternalServerError)
                w.Write([]byte(`{"error": "Internal Server Error"}`))
            }
        }()
        next.ServeHTTP(w, r)
    })
}