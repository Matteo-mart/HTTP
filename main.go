package main

import (
	"fmt"
	"log"
	"my_srv_http/recordings"
	"net/http"
)

func HandlerHello2GET(response http.ResponseWriter, request *http.Request) {
	q := request.URL.Query()
	response.Write([]byte(fmt.Sprintf("GET hello %s", q.Get("name"))))
}

func HandlerHello2POST(response http.ResponseWriter, request *http.Request) {
	q := request.URL.Query()
	response.Write([]byte(fmt.Sprintf("POST hello %s", q.Get("name"))))
}

func HandleGETlist(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte(fmt.Sprintf("todolist 1, todolist 2")))
}

func Handletest(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte(fmt.Sprintf(
		"test scheme: %s, host: %s, port: %s, path: %s, bucket: %s, objectname: %s",
		request.URL.Scheme,
		request.Host,
		request.URL.Port(),
		request.URL.Path,
		request.PathValue("bucket"),
		request.PathValue("objectname"),
	)))
}

func HandleTDlist(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte(fmt.Sprintf(
		"to do list ID: %s",
		request.PathValue("ID"),
		//request.PathValue("to do list"),
	)))
}

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

	mux.HandleFunc("GET /hello2", HandlerHello2GET)
	mux.HandleFunc("POST /hello2", HandlerHello2POST)
	mux.HandleFunc("/b/{bucket}/o/{objectname...}", Handletest)
	mux.HandleFunc("/c/{bucket}/o/{objectname}", Handletest)

	mux.HandleFunc("GET /list", HandleGETlist)
	mux.HandleFunc("GET /list/{ID}", HandleTDlist)

	mux.HandleFunc("GET /recordings/list/{ID}", recordings.HandleGETRecording)
	mux.HandleFunc("GET /recordings/list", recordings.HandleGETlist)

	http.ListenAndServe(":3000", mux)
}
