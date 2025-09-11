package tests

import (
	"fmt"
	"net/http"
)

func HandleTDlist(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte(fmt.Sprintf(
		"to do list ID: %s",
		request.PathValue("ID"),
		//request.PathValue("to do list"),
	)))
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

func HandlerHello2GET(response http.ResponseWriter, request *http.Request) {
	q := request.URL.Query()
	response.Write([]byte(fmt.Sprintf("GET hello %s", q.Get("name"))))
}

func HandlerHello2POST(response http.ResponseWriter, request *http.Request) {
	q := request.URL.Query()
	response.Write([]byte(fmt.Sprintf("POST hello %s", q.Get("name"))))
}
