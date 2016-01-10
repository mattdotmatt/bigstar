package models

type Character struct {
	FirstName string `json:"firstName" validate:"nonzero"`
	LastName  string `json:"lastName"`
}
