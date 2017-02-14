package models

type SelectionOption struct {
	Attribute string   `json:"attribute"`
	Enum      []string `json:"enum,omitempty"`
	Required  bool     `json:"required"`
	Type      string   `json:"type"`
	Metric    string   `json:"metric,omitempty"`
}

type SelectionOptionModel struct{}

func (sel *SelectionOptionModel) FindAll() []SelectionOption {
	return []SelectionOption{
		{Attribute: "Gender", Enum: []string{"male", "female"}, Required: true, Type: "String"},
		{Attribute: "Ambience", Enum: []string{"Conf Room", "Office", "Car", "Street"}, Required: true, Type: "String"},
		{Attribute: "Accent", Enum: []string{"English", "Indian", "American", "Hispanic", "Chinese"}, Required: true, Type: "String"},
		{Attribute: "Speech Input", Enum: []string{"Headphone", "Phone Mic"}, Required: true, Type: "String"},
		{Attribute: "Distance", Metric: "cm", Required: true, Type: "Integer"},
		{Attribute: "Audio Sampling Rate", Metric: "hz", Required: true, Type: "Integer"},
		{Attribute: "Speech Type", Enum: []string{"Command", "Sentence", "Paragraph"}, Required: true, Type: "String"},
	}
}
