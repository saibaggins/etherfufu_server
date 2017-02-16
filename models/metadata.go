package models

import (
	"github.com/saibaggins/etherfufu-server/core"
	"github.com/saibaggins/etherfufu-server/utils"
)

var (
	MetadataModelTableName = core.ActiveENV() + "_" + "etherfufu_audio_sample_metadata"
)

type Metadata struct {
	ID          string `json:"id,omitempty"`
	Gender      string `json:"gender" binding:"required"`
	Ambience    string `json:"ambience" binding:"required"`
	Accent      string `json:"accent" binding:"required"`
	SpeechInput string `json:"speech_input" binding:"required"`
	Distance    int    `json:"distance" binding:"required"`
	SampleRate  int    `json:"sample_rate" binding:"required"`
	SpeechType  string `json:"speech_type" binding:"required"`
}

type MetadataModel struct{}

func (self *MetadataModel) Put(data Metadata) bool {
	model := utils.NewDBModel(MetadataModelTableName)
	return model.PutRecord(data)
}
