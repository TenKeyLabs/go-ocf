package exercise

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/exercise"
)

// Object describing equity compensation exercise transaction
type EquityCompensationExercise struct {
	exercise.Exercise
	transactions.SecurityTransaction

	// Quantity of shares exercised
	Quantity string `json:"quantity"`
}
