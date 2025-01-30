package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Entry represents a calorie entry in the database
type Entry struct {
	ID          primitive.ObjectID `bson:"id"`
	Dish        *string            `json:"dish"`
	Fat         *float64           `json:"fat"`
	Ingredients *string            `json:"ingredients"`
	Calories    *string            `json:"calories"`
}