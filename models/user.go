package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"-"`
	UserID      int           `bson:"userId" json:"userId"`
	Username    string        `bson:"username" json:"username"`
	RatingColor string        `bson:"ratingColor" json:"ratingColor"`
	RatingTitle string        `bson:"ratingTitle" json:"ratingTitle"`
	HasPower    bool          `bson:"hasPower" json:"hasPower"`
}
