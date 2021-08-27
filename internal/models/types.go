package models

type Dog struct {
	Name   string `json:"Name" validate:"required"`
	Gender string `json:"Gender"`
}
