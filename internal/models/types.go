package models

// Dog struct for dogs.
type Dog struct {
	Name   string `json:"name" validate:"required"`
	Gender string `json:"gender" validate:"required"`
}
