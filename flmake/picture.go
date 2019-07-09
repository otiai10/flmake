package flmake

import (
	"bytes"

	"gopkg.in/yaml.v2"
)

// Picture is one of Elements, representing an image scrap which can be drawn on base image.
type Picture struct {
	Image    Image    `yaml:"image"`
	Position Position `yaml:"position"`
	Scale    float64  `yaml:"scale"`
}

// Decode ...
// TODO: Ugly
func (pict *Picture) Decode(e map[string]interface{}) error {
	buf := bytes.NewBuffer(nil)
	if err := yaml.NewEncoder(buf).Encode(e); err != nil {
		return err
	}
	if err := yaml.NewDecoder(buf).Decode(pict); err != nil {
		return err
	}
	return nil
}
