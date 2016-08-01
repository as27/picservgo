package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GETFolderElements returns the folder content.
func GETFolderElements(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	folderPath := vars["folder"]
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		log.Println(err)
	}
	fc := NewFolderContent(folderPath, &MyGallery)
	for _, f := range files {
		fc.Add(f)
	}
	b, err := json.Marshal(fc)
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write(b)
	if err != nil {
		log.Fatal(err)
	}

}