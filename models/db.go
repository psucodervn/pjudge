package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
)

type databaseConnection struct {
	Host     string
	Database string
	Username string
	Password string
}

var conn *databaseConnection
var session *mgo.Session

func init() {
	conn = &databaseConnection{
		Host:     beego.AppConfig.DefaultString("DB.Host", "localhost:27017"),
		Database: beego.AppConfig.DefaultString("DB.Database", "pjudge"),
		Username: beego.AppConfig.DefaultString("DB.Username", ""),
		Password: beego.AppConfig.DefaultString("DB.Password", ""),
	}
	session = createSession(conn)
}

func createSession(conn *databaseConnection) *mgo.Session {
	dialInfo := &mgo.DialInfo{
		Addrs:    strings.Split(conn.Host, ","),
		Timeout:  60 * time.Second,
		Database: conn.Database,
		Username: conn.Username,
		Password: conn.Password,
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		beego.Error("Connect to " + conn.Host + "/" + conn.Database + " failed!")
		beego.Critical(err.Error())
		return nil
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}

func Count(collection string, query *bson.M) (count int, err error) {
	sess := session.Copy()
	defer sess.Close()
	coll := sess.DB(conn.Database).C(collection)

	count, err = coll.Find(query).Count()
	return
}

func query(coll *mgo.Collection, query *bson.M, fields *[]string, sort string, skip, limit int) *mgo.Query {
	q := coll.Find(query)
	if fields != nil && len(*fields) > 0 {
		selector := bson.M{}
		for _, s := range *fields {
			selector[s] = 1
		}
		q = q.Select(selector)
	}
	if skip > 0 {
		q = q.Skip(skip)
	}
	if limit > 0 {
		q.Limit(limit)
	}
	if len(sort) > 0 {
		q.Sort(sort)
	}
	return q
}
