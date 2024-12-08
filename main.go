package main

import (
	"JsonLogs/service"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(service.FrontInit))
	http.Handle("POST /sendLog", http.HandlerFunc(service.SendLogHandler))
	http.Handle("/view-logs", http.HandlerFunc(service.LogViewHandler))

	http.ListenAndServe(":8080", nil)
}
