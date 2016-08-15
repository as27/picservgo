package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var ServerPort string

var GalleryRootFolder string

var MyGallery = Imaginary{}

func init() {
	LoadConf()

	ServerPort = Conf.ServerPort

	GalleryRootFolder = Conf.GalleryRoot

	MyGallery = Imaginary{
		URL:         Conf.Imaginary.URL,
		ThumbWidth:  Conf.Imaginary.ThumbWidth,
		ThumbHeight: Conf.Imaginary.ThumbHeight,
		ThumbMethod: Conf.Imaginary.ThumbMethod,
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc(`/folder/{folder:.*}`, GETFolderElements).Methods("GET")
	router.PathPrefix(`/static/`).Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(Conf.HtmlRoot))))
	//log.Println("Starting server at port " + Conf.ServerPort)
	log.Fatal(http.ListenAndServe(ServerPort, router))
}
