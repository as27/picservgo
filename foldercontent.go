package main

import (
	"os"
	"path"
	"strings"
)

// FolderContent represents the content of a folder
type FolderContent struct {
	FolderPath string         `json:"path"`    //The path to the folder
	Subfolders []string       `json:"folders"` //All the subfolders
	Images     []BlueimpImage `json:"images"`  //Just the images
	Gallery    Gallery        `json:"-"`
}

//NewFolderContent returns a pointer to a FolderContent instance
func NewFolderContent(folderPath string, g Gallery) *FolderContent {
	return &FolderContent{
		FolderPath: folderPath,
		Gallery:    g,
	}
}

//Add adds a new element to the FolderContent
func (fc *FolderContent) Add(fi os.FileInfo) {
	if fi.IsDir() {
		fc.Subfolders = append(fc.Subfolders, fi.Name())
	}
	if hasPicSuffix(fi.Name()) {
		fpath := path.Join(fc.FolderPath, fi.Name())
		img := BlueimpImage{
			Title: fi.Name(),
			Href:  fc.Gallery.FullSizeURL(fpath),
			Type:  "image/jpeg",
			Thumb: fc.Gallery.ThumbURL(fpath),
		}
		fc.Images = append(fc.Images, img)
	}

}

//PicSuffixes all supported picture suffixes in lower case
var PicSuffixes = []string{
	".jpg",
	".jpeg",
	".gif",
	".png",
}

func hasPicSuffix(name string) bool {
	ext := strings.ToLower(path.Ext(name))
	for _, ps := range PicSuffixes {
		if ps == ext {
			return true
		}
	}
	return false
}

// BlueimpImage is the implementation for the Blueimp Gallery
// https://blueimp.github.io/Gallery/
type BlueimpImage struct {
	Title string `json:"title"`
	Href  string `json:"href"`
	Type  string `json:"type"`
	Thumb string `json:"thumbnail"`
}

// MarshalJSON is the implementation of the json.Marshaler interface
/*func (i *BlueimpImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(i)
}*/
