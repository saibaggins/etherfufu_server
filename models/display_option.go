package models

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/saibaggins/etherfufu-server/core"
	"github.com/saibaggins/etherfufu-server/utils"
)

type DisplayOption struct {
	Attribute   string   `json:"attribute"`
	DisplayName string   `json:"display_name"`
	Enum        []string `json:"enum,omitempty"`
	Required    bool     `json:"required"`
	Type        string   `json:"type"`
	Metric      string   `json:"metric,omitempty"`
}

var (
	DisplayOptionModelTableName = core.ActiveENV() + "_" + "etherfufu_display_options"
)

type DisplayOptionModel struct{}

func (self *DisplayOptionModel) FindAll() []DisplayOption {
	var display_options []DisplayOption
	table := utils.NewDBModel(DisplayOptionModelTableName)
	result := table.ListRecords()
	dynamodbattribute.UnmarshalListOfMaps(result.Items, &display_options)
	return display_options
}

func (self *DisplayOptionModel) Put(options []DisplayOption) bool {
	model := utils.NewDBModel(DisplayOptionModelTableName)
	var success bool
	for _, option := range options {
		success = model.PutRecord(option)
	}

	return success
}
