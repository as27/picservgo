package main

import (
	"flag"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// ConfFile is the path to the configuration file
var ConfFile = flag.String("conf", "picservgo.yaml", "Path to the conf (yaml) file")

// Conf stores the configuration
var Conf Options

// Options is the structure of the config file
type Options struct {
	ServerPort  string    `yaml:"ServerPort"`
	GalleryRoot string    `yaml:"GalleryRoot"`
	HtmlRoot    string    `yaml:"HtmlRoot"`
	Imaginary   Imaginary `yaml:"Imaginary"`
}

type dependencies struct{}

func init() {
	flag.Parse()
}

func LoadConf() {
	b, err := ioutil.ReadFile(*ConfFile)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(b, &Conf)
	if err != nil {
		log.Fatal(err)
	}
}
