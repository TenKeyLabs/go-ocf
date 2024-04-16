package exercise

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/exercise"
)

// Object describing equity compensation exercise transaction
type EquityCompensationExercise struct {
	exercise.Exercise
	transactions.Transaction
	transactions.SecurityTransaction
	objects.Object

	// Quantity of shares exercised
	Quantity string `json:"quantity"`
}
