package main

import (
	"log"
	"my_srv_http/recordings"
	"my_srv_http/tests"
	"net/http"
)

func main() {

	log.Printf("on va se connecter a la base de donnée")
	err := recordings.InitDb()
	if err != nil {
		log.Printf("Impossible de se connecter à la base de donnée, err: %s", err)
		return
	}

	log.Printf("configuration du serveur http")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello2", tests.HandlerHello2GET)
	mux.HandleFunc("POST /hello2", tests.HandlerHello2POST)
	mux.HandleFunc("/b/{bucket}/o/{objectname...}", tests.Handletest)
	mux.HandleFunc("/c/{bucket}/o/{objectname}", tests.Handletest)

	mux.HandleFunc("GET /list", tests.HandleGETlist)
	mux.HandleFunc("GET /list/{ID}", tests.HandleTDlist)

	mux.HandleFunc("GET /recordings/list/{ID}", recordings.HandleGETRecording)
	mux.HandleFunc("GET /recordings/list", recordings.HandleGETlist)

	log.Printf("listening ..")
	http.ListenAndServe(":3000", mux)
}
