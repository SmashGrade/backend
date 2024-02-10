package dao

import (
	"testing"

	"github.com/SmashGrade/backend/app/models"
)

// Check if a slice can be asserted correctly and keep all data intact
func TestAssertSlice(t *testing.T) {
	inputSlice := make([]any, 0)

	inputSlice = append(inputSlice, models.Module{
		Description: "Test01",
	})

	inputSlice = append(inputSlice, models.Module{
		Description: "Supertest",
	})

	outputSlice := assertSlice[models.Module](inputSlice)

	if len(outputSlice) != len(inputSlice) {
		t.Fatalf("expected slice len %v got %v", len(inputSlice), len(outputSlice))
	}

	for i := range outputSlice {
		inputModule, assertionOk := inputSlice[i].(models.Module)
		if assertionOk == false {
			t.Fatalf("can not assert input slice as module")
		}
		if outputSlice[i].Description != inputModule.Description {
			t.Fatalf("input and output description differ")
		}
	}
}
