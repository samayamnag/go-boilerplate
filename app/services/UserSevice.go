package services

import (
	"fmt"
	"time"
	"github.com/janaonline/icmyc/config"
	"github.com/janaonline/icmyc/app/models"
	"gopkg.in/mgo.v2/bson"
	"github.com/janaonline/icmyc/app/util"
	"github.com/janaonline/icmyc/app/database"
)

const (
	USER_COLLECTION = "users"
)

type Repository struct{}

func InsertUser(user models.User) models.User{
	session, err := database.MongoConnection()

	if err != nil {
		panic(err.Error())
	}

	defer session.Close()

	now := time.Now()
	user.ID = bson.NewObjectId()
	user.Password = util.HashAndSalt(user.Password)
	user.CreatedAt = now
	user.BeforeInsert()

	session.DB(config.Get().MongoDBName).C(USER_COLLECTION).Insert(user)

	return user
}

func GetAllUsers() models.Users {
	session, err := database.MongoConnection()
	if err != nil {
		panic(err.Error())
	}

	defer session.Close()

	users := models.Users{}

	c := session.DB(config.Get().MongoDBName).C(USER_COLLECTION)

	if err := c.Find(nil).All(&users); err != nil {
		panic(err.Error())
	}

	return users
}

func FindById(id string) models.User {
	session, err := database.MongoConnection()
	
	if err != nil {
		panic(err.Error())
	}

	defer session.Close()

	var user models.User

	c := session.DB(config.Get().MongoDBName).C(USER_COLLECTION)

	if err = c.FindId(bson.ObjectIdHex(id)).One(&user); err != nil {
		fmt.Println("Failed to write result:", err)
	}
	return user
}