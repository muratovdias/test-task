package models

type User struct {
	Email    string `json:"email" bson:"email, omitempty"`
	Salt     string `json:"salt" bson:"salt, omitempty"`
	Password string `json:"password" bson:"password, omitempty"`
}
