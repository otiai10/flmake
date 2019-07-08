package flmake

import "fmt"

// Config ...
type Config struct {
	Image  Image   `yaml:"image"`
	Images []Image `yaml:"images"`
	Labels []Label `yaml:"labels"`
}

// Populate ...
func (c *Config) Populate() error {
	if len(c.Images) == 0 {
		if len(c.Image) == 0 {
			return fmt.Errorf("`image` or `images` fields are required to make a flyer")
		}
		c.Images = []Image{c.Image}
	}
	return nil
}
