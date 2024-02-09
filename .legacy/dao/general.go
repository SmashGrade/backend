package dao

import (
	"encoding/json"
)

func ParseEntityToSchema[T any, U any](fromEntity T, fromSchema U) error {

	jsonData, err := json.Marshal(&fromEntity)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, &fromSchema)
	if err != nil {
		return err
	}

	return nil
}

func ParseSchemaToEntity[T any, U any](fromSchema T, fromEntity U) error {

	jsonData, err := json.Marshal(&fromSchema)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, &fromEntity)
	if err != nil {
		return err
	}

	return nil
}

func findMax(arr []uint) uint {
	if len(arr) == 0 {
		return 0
	}

	max := arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}

	return max
}