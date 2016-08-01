package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var ServerPort = ":4527"

var MyGallery = Imaginary{
	URL:         Conf.Imaginary.URL,
	ThumbWidth:  Conf.Imaginary.ThumbWidth,
	ThumbHeight: Conf.Imaginary.ThumbHeight,
	ThumbMethod: Conf.Imaginary.ThumbMethod,
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc(`/folder/{folder}`, GETFolderElements).Methods("GET")
	//log.Println("Starting server at port " + Conf.ServerPort)
	log.Fatal(http.ListenAndServe(ServerPort, router))
}
