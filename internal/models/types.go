package models

// Dog struct for dogs.
type Dog struct {
	Name   string `json:"name" validate:"required"`
	Gender string `json:"gender" validate:"required"`
}

// NewDog creates new dog object.
func NewDog(name string, gender string) Dog {
	return Dog{
		Name:   name,
		Gender: gender,
	}
}
