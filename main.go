package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/otiai10/flmake/flmake"
	"gopkg.in/yaml.v2"
)

var (
	configpath string
)

func init() {
	flag.StringVar(&configpath, "c", "./config.yml", "Path to config file.")
	flag.Parse()
}

func main() {
	fmt.Println(configpath)
	c, err := os.Open(configpath)
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	config := new(flmake.Config)
	if err := yaml.NewDecoder(c).Decode(config); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v\n", config)
}
