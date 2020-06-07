package main

import "net/http"
import "timewheel/Handler"

func main() {

	httpHandler := &Handler.HttpHandler{}
	http.HandleFunc("/", httpHandler.AddTimeWheel)
	http.ListenAndServe("127.0.0.1:8000", nil)
}
