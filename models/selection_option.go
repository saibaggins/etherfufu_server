package models

type SelectionOption struct {
	Attribute string   `json:"attribute"`
	Values    []string `json:"values"`
	Metric    string   `json:"metric,omitempty"`
}

type SelectionOptionModel struct{}

func (sel *SelectionOptionModel) FindAll() []SelectionOption {
	return []SelectionOption{
		SelectionOption{"Name", []string{"Hello", "World"}, ""},
		SelectionOption{"Recording Type", []string{"mic", "headphone"}, ""},
	}
}
