package recordings

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleGETlist(response http.ResponseWriter, request *http.Request) {

	albums, err := GetAlbums()
	if err != nil {
		// ca plante
		msg := fmt.Sprintf("il y a une erreur: %s", err)
		response.WriteHeader(500)
		response.Write([]byte(msg))
		return
	}

	jsonStr, err := json.Marshal(albums)
	if err != nil {
		// ca plante
		msg := fmt.Sprintf("il y a une erreur: %s", err)
		response.WriteHeader(500)
		response.Write([]byte(msg))
		return
	}
	response.Header().Add("content-type", "application/json")
	response.WriteHeader(200)
	response.Write([]byte(jsonStr))

}

func HandleGETRecording(response http.ResponseWriter, request *http.Request) {

	id := request.PathValue("ID")

	album, err := GetAlbum(id)
	if err != nil {
		msg := fmt.Sprintf("il y a une erreur: %s", err)
		response.WriteHeader(500)
		response.Write([]byte(msg))
		return
	}

	// response.Write([]byte(fmt.Sprintf(
	// 	"recording ID: %s",
	// 	id)))

	jsonStr, err := json.Marshal(album)

	if err != nil {
		// ca plante
		msg := fmt.Sprintf("il y a une erreur: %s", err)
		response.WriteHeader(500)
		response.Write([]byte(msg))
		return
	}
	response.Header().Add("content-type", "application/json")
	response.WriteHeader(200)
	response.Write([]byte(jsonStr))

}
