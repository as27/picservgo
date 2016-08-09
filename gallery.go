package main

import (
	"log"
	"net/url"
	"path/filepath"
)

// Gallery abstracts the methods which are needed for a gallery
type Gallery interface {
	FullSizeURL(string) string
	ThumbURL(string) string
}

// Imaginary is the Gallery implementation when the imaginary microservice
// is used.
// The yaml tag definition is for the configuration file
type Imaginary struct {
	URL         string `yaml:"url"` // Url to the microservice
	ThumbWidth  string `yaml:"thumbWidth"`
	ThumbHeight string `yaml:"thumbHeight"`
	ThumbMethod string `yaml:"thumbMethod"`
}

// FullSizeURL returns the url for the full size image
func (i *Imaginary) FullSizeURL(fp string) string {
	fp = filepath.ToSlash(fp)
	u, err := url.Parse(i.URL + "/fullsize")
	if err != nil {
		log.Println(err)
	}
	q := u.Query()
	q.Set("file", fp)
	u.RawQuery = q.Encode()
	return u.String()
}

// ThumbURL returns the url to the thumb image
func (i *Imaginary) ThumbURL(fp string) string {
	fp = filepath.ToSlash(fp)
	u, err := url.Parse(i.URL + "/" + i.ThumbMethod)
	if err != nil {
		log.Println(err)
	}
	q := u.Query()
	q.Set("file", fp)
	q.Set("width", i.ThumbWidth)
	q.Set("height", i.ThumbHeight)
	u.RawQuery = q.Encode()
	return u.String()
}
