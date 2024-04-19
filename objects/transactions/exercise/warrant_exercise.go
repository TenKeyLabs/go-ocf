package exercise

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/exercise"
)

// Object describing a warrant exercise transaction
type WarrantExercise struct {
	exercise.Exercise
	transactions.SecurityTransaction

	// What is the id of the warrant's exercise trigger that resulted in this exercise
	TriggerID string `json:"trigger_id"`
}
