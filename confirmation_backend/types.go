package main

type Person struct {
	UserId string `json:"UserId,omitempty" bson:"UserId,omitempty"`
	Name   string `json:"name,omitempty" bson:"name,omitempty"`
	Email  string `json:"email,omitempty" bson:"email,omitempty"`
	Mobile string `json:"mobile,omitempty" bson:"mobile,omitempty"`
}
