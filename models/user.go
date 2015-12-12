package models

import (
	"github.com/psufighter/pjudge/conf"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID             bson.ObjectId `bson:"_id,omitempty" json:"-"`
	UserID         int           `bson:"userId" json:"userId"`
	Username       string        `bson:"username" json:"username"`
	Name           string        `bson:"name" json:"name"`
	HashedPassword string        `bson:"hashedPassword" json:"-"`
	Email          string        `bson:"email" json:"email"`
	Privilege      int           `bson:"privilege" json:"privilege"`
	RatingColor    string        `bson:"ratingColor" json:"ratingColor"`
	RatingTitle    string        `bson:"ratingTitle" json:"ratingTitle"`
	HasPower       bool          `bson:"hasPower" json:"hasPower"`
	RegisterTime   int64         `bson:"registerTime" json:"registerTime"`
}

func UserByUserName(username string, user *User) (err error) {
	sess := session.Copy()
	defer sess.Close()
	coll := sess.DB(conn.Database).C(conf.CUser)

	err = coll.Find(bson.M{"username": username}).One(user)
	return
}

func UserExisted(username string) bool {
	var user User
	return UserByUserName(username, &user) != mgo.ErrNotFound
}

func UserInsert(user *User) (err error) {
	sess := session.Copy()
	defer sess.Close()
	coll := sess.DB(conn.Database).C(conf.CUser)

	return coll.Insert(user)
}
