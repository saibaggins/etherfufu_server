package models

type DisplayOption struct {
	Attribute string   `json:"attribute"`
	DisplayName string `json:"display_name"`
	Enum      []string `json:"enum,omitempty"`
	Required  bool     `json:"required"`
	Type      string   `json:"type"`
	Metric    string   `json:"metric,omitempty"`
}

type DisplayOptionModel struct{}

func (sel *DisplayOptionModel) FindAll() []DisplayOption {
	return []DisplayOption{
		{Attribute: "gender", DisplayName: "Gender", Enum: []string{"male", "female"}, Required: true, Type: "String"},
		{Attribute: "ambience", DisplayName: "Ambience", Enum: []string{"Conf Room", "Office", "Car", "Street"}, Required: true, Type: "String"},
		{Attribute: "accent", DisplayName: "Accent", Enum: []string{"English", "Indian", "American", "Hispanic", "Chinese"}, Required: true, Type: "String"},
		{Attribute: "speech_input", DisplayName: "Speech Input", Enum: []string{"Headphone", "Phone Mic"}, Required: true, Type: "String"},
		{Attribute: "distance", DisplayName: "Distance", Metric: "cm", Required: true, Type: "Integer"},
		{Attribute: "sample_rate", DisplayName: "Audio Sampling Rate", Metric: "hz", Required: true, Type: "Integer"},
		{Attribute: "speech_type", DisplayName: "Speech Type", Enum: []string{"Command", "Sentence", "Paragraph"}, Required: true, Type: "String"},
	}
}
