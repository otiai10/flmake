package flmake

import (
	"bytes"

	"gopkg.in/yaml.v2"
)

// Label is one of Elements, representing a text drawn on base mage.
type Label struct {
	Text     string   `yaml:"text"`
	Position Position `yaml:"position"`
	Size     float64  `yaml:"size"`
}

// Decode ...
func (label *Label) Decode(e map[string]interface{}) error {
	buf := bytes.NewBuffer(nil)
	if err := yaml.NewEncoder(buf).Encode(e); err != nil {
		return err
	}
	if err := yaml.NewDecoder(buf).Decode(label); err != nil {
		return err
	}
	return nil

}
