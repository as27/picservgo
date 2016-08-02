package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// ConfFile is the path to the configuration file
var ConfFile = "picservgo.yaml"

// Conf stores the configuration
var Conf Options

// Options is the structure of the config file
type Options struct {
	ServerPort  string    `yaml:"ServerPort"`
	GalleryRoot string    `yaml:"GalleryRoot"`
	Imaginary   Imaginary `yaml:"Imaginary"`
}

type dependencies struct{}

func LoadConf() {
	b, err := ioutil.ReadFile(ConfFile)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(b, &Conf)
	if err != nil {
		log.Fatal(err)
	}
}
