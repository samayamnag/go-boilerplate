package database

import (
	"fmt"
	"time"
	"gopkg.in/mgo.v2"

	"github.com/janaonline/icmyc/config"
)

func MongoConnection() (mongoSession *mgo.Session, err error){
	connStr := fmt.Sprintf("%s:%d", config.Get().MongoDBHost, config.Get().MongoDBPort)
	// We need this object to establish a session to our MongoDB.
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{connStr},
		Timeout:  60 * time.Second,
		Database: config.Get().MongoDBName,
		Username: config.Get().MongoDBUsername,
		Password: config.Get().MongoDBPassword,
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	mongoSession, err = mgo.DialWithInfo(mongoDBDialInfo)
	return 
}