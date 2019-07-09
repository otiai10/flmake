package flmake

import (
	"fmt"
	"path/filepath"
)

// Config ...
type Config struct {
	Name     string
	Image    Image    `yaml:"image"`
	Images   []Image  `yaml:"images"`
	Elements Elements `yaml:"elements"`
}

// Populate ...
func (c *Config) Populate() error {
	if len(c.Images) == 0 {
		if len(c.Image) == 0 {
			return fmt.Errorf("`image` or `images` fields are required to make a flyer")
		}
		c.Images = []Image{c.Image}
	}
	dir := filepath.Dir(c.Name)
	for i, img := range c.Images {
		c.Images[i] = Image(filepath.Join(dir, string(img)))
	}
	return nil
}

// Override ...
func (c *Config) Override(base *Config) error {
	return nil
}
