package utils

import (
	"gopkg.in/mgo.v2"
)

type DBConfig struct {
	URL  string `default:"mongo://localhost:3600/test"`
	Port uint   `default:"3306"`
}

var db *mgo.Session

// return a DB session ...
func StartDBSession(url string) {
	dialInfo, err := mgo.ParseURL(url)
	if err != nil {
		panic(err)
	}

	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	db = session
}

func GetDB() *mgo.Session {
	return db
}