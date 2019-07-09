package flmake

// Elements ...
type Elements []Element

// UnmarshalYAML ...
func (elms *Elements) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var raw []map[string]interface{}
	if err := unmarshal(&raw); err != nil {
		return err
	}
	var values Elements = []Element{}
	for _, e := range raw {
		if _, ok := e["text"]; ok {
			label := new(Label)
			if err := label.Decode(e); err != nil {
				return err
			}
			values = append(values, label)
		} else if _, ok := e["image"]; ok {
			pict := new(Picture)
			if err := pict.Decode(e); err != nil {
				return err
			}
			values = append(values, pict)
		}
	}
	*elms = values
	return nil
}

// Element represents anything which can be drawn on a base image.
type Element interface{}
