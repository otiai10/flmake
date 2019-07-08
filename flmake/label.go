package flmake

// Label ...
type Label struct {
	Text     string   `yaml:"text"`
	Position Position `yaml:"position"`
	Size     float64  `yaml:"size"`
}
