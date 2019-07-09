package main

import (
	_ "image/jpeg"

	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

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
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	c, err := os.Open(configpath)
	if err != nil {
		return err
	}
	defer c.Close()

	abspath, _ := filepath.Abs(c.Name())
	config := &flmake.Config{Name: abspath}

	b, err := ioutil.ReadAll(c)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(b, config); err != nil {
		return err
	}

	if err := config.Populate(); err != nil {
		return err
	}

	builder := &flmake.Builder{Config: config}
	if err := builder.Build(); err != nil {
		return err
	}

	return nil
}
