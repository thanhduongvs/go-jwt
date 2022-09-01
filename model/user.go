package model

import (
	fake "github.com/brianvoe/gofakeit/v6"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name"`
	Email  string             `json:"email" bson:"email"`
	Gender string             `json:"gender"`
	Phone  string             `json:"phone"`
	Age    int64              `json:"age"`
	UserId string             `json:"user_id"`
}

func NewUser() User {
	id := primitive.NewObjectID()

	name := fake.Name()
	email := fake.Email()
	phone := fake.Phone()
	gender := fake.Gender()
	age := int64(fake.Number(10, 99))
	user := User{
		ID:     id,
		Name:   name,
		Email:  email,
		Gender: gender,
		Phone:  phone,
		Age:    age,
		UserId: id.Hex(),
	}
	return user
}
