package models

import "github.com/satori/go.uuid"

type Metadata struct {
	ID uuid.UUID `json:"id,omitempty"`
	Gender string `json:"gender" binding:"required"`
	Ambience string `json:"ambience" binding:"required"`
	Accent string `json:"accent" binding:"required"`
	SpeechInput string `json:"speech_input" binding:"required"`
	Distance int	`json:"distance" binding:"required"`
	SampleRate int `json:"sample_rate" binding:"required"`
	SpeechType string `json:"speech_type" binding:"required"`
}
