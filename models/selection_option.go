package models

type SelectionOption struct {
	Attribute string   `json:"attribute"`
	Enum      []string `json:"enum,omitempty"`
	Required  bool     `json:"required"`
	Metric    string   `json:"metric,omitempty"`
}

type SelectionOptionModel struct{}

func (sel *SelectionOptionModel) FindAll() []SelectionOption {
	return []SelectionOption{
		{Attribute: "Accent", Enum: []string{"English", "Indian", "American", "Hispanic", "Chinese"}, Required: true},
		{Attribute: "Recording Type", Enum: []string{"mic", "headphone"}, Required: true},
		{Attribute: "Height", Metric: "cm", Required: true},
		{Attribute: "Audio Sample Frequency", Metric: "hz", Required: true},
	}
}
