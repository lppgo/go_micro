package db

import (
	"time"

	"github.com/globalsign/mgo"
)

func DialMongo(addr string) *mgo.Session {
	sess, err := mgo.DialWithTimeout(addr, 5*time.Second)
	if err != nil {
		panic(err)
	}
	sess.SetMode(mgo.Monotonic, true)
	err = sess.Ping()
	if err != nil {
		panic(err)
	}
	return sess
}
