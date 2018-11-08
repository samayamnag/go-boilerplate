package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID			bson.ObjectId	`bson:"_id"`
	Email		string			`bson:"email"`
	Password	string			`bson:"password"`
	CreatedAt	time.Time		`bson:"created_at"`
	UpdatedAt	time.Time		`bson:"updated_at"`
}

type Users []User

func(user *User) BeforeInsert() error{
	user.UpdatedAt = time.Now()
	return nil
}

func (user *User) BeforeUpdate() error {
	user.UpdatedAt = time.Now()
	return nil
}