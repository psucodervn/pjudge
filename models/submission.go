package models

import (
	"github.com/psufighter/pjudge/conf"
	"gopkg.in/mgo.v2/bson"
)

type Submission struct {
	ID                  bson.ObjectId `bson:"_id,omitempty" json:"-"`
	CompileErrorMessage string        `bson:"compileErrorMessage" json:"compileErrorMessage"`
	Language            string        `bson:"language" json:"language"`
	LanguageID          int           `bson:"languageId" json:"languageId"`
	Memory              int           `bson:"memory" json:"memory"`
	ProblemID           int           `bson:"problemId" json:"problemId"`
	ProblemName         string        `bson:"problemName" json:"problemName"`
	Rating              int           `bson:"rating" json:"rating"`
	RatingColor         string        `bson:"ratingColor" json:"ratingColor"`
	Runtime             int           `bson:"runtime" json:"runtime"`
	Code                string        `bson:"code" json:"solution"`
	SubmissionID        int           `bson:"submissionId" json:"submissionId"`
	SubmitTime          int64         `bson:"submitTime" json:"submitTime"`
	UserID              int           `bson:"userId" json:"userId"`
	Username            string        `bson:"username" json:"username"`
	Verdict             string        `bson:"verdict" json:"verdict"`
	VerdictColor        string        `bson:"verdictColor" json:"verdictColor"`
	VerdictID           int           `bson:"verdictId" json:"verdictId"`
}

func SubmissionAll(query *bson.M, subs *[]Submission) (err error) {
	sess := session.Copy()
	defer sess.Close()
	coll := sess.DB(conn.Database).C(conf.CSubmission)

	return coll.Find(query).All(subs)
}

func SubmissionList(query *bson.M, subs *[]Submission, fields *[]string, sort *string, skip, limit *int) (err error) {
	sess := session.Copy()
	defer sess.Close()
	coll := sess.DB(conn.Database).C(conf.CSubmission)

	return makeQuery(coll, query, fields, sort, skip, limit).All(subs)
}

func SubmissionFind(query *bson.M, sub *Submission) (err error) {
	sess := session.Copy()
	defer sess.Close()
	coll := sess.DB(conn.Database).C(conf.CSubmission)

	return coll.Find(query).One(sub)
}

func SubmissionInsert(sub *Submission) (err error) {
	sess := session.Copy()
	defer sess.Close()
	coll := sess.DB(conn.Database).C(conf.CSubmission)

	return coll.Insert(sub)
}
