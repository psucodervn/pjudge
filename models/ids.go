package models

import (
	"github.com/psufighter/pjudge/conf"
	"gopkg.in/mgo.v2/bson"
)

type IDs struct {
	Name string `bson:"name"`
	ID   int    `bson:"id"`
}

func LastID(name string, id *int) (err error) {
	sess := session.Copy()
	defer sess.Close()
	coll := sess.DB(conn.Database).C(conf.CIDs)

	var ids IDs
	err = coll.Find(bson.M{"name": name}).One(&ids)
	if err != nil {
		return
	}

	*id = ids.ID
	return
}

func IncreaseID(name string) (err error) {
	sess := session.Copy()
	defer sess.Close()
	coll := sess.DB(conn.Database).C(conf.CIDs)

	_, err = coll.Upsert(bson.M{"name": name}, bson.M{"$inc": bson.M{"id": 1}})
	return
}
