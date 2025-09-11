package main

import (
	"log"
	"my_srv_http/recordings"
	"my_srv_http/tests"
	"net/http"
)

func main() {
	log.Printf("coucou")
	mux := http.NewServeMux()

	// redirection
	rh := http.RedirectHandler("http://example.org", http.StatusTemporaryRedirect)
	mux.Handle("/foo", rh)

	// texte
	mux.HandleFunc("/hello", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("hello " + request.Method))
	})

	mux.HandleFunc("/test", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("test " + request.Method))
	})

	mux.HandleFunc("GET /hello2", tests.HandlerHello2GET)
	mux.HandleFunc("POST /hello2", tests.HandlerHello2POST)
	mux.HandleFunc("/b/{bucket}/o/{objectname...}", tests.Handletest)
	mux.HandleFunc("/c/{bucket}/o/{objectname}", tests.Handletest)

	mux.HandleFunc("GET /list", tests.HandleGETlist)
	mux.HandleFunc("GET /list/{ID}", tests.HandleTDlist)

	mux.HandleFunc("GET /recordings/list/{ID}", recordings.HandleGETRecording)
	mux.HandleFunc("GET /recordings/list", recordings.HandleGETlist)

	http.ListenAndServe(":3000", mux)
}
