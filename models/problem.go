package models

import (
	"github.com/psufighter/pjudge/conf"
	"gopkg.in/mgo.v2/bson"
)

type PDetail struct {
	ID                            bson.ObjectId          `bson:"_id,omitempty" json:"-"`
	Author                        string                 `bson:"author" json:"author"`
	DefaultCheckerExactMatch      bool                   `bson:"defaultCheckerExactMatch" json:"defaultCheckerExactMatch"`
	Difficulty                    float64                `bson:"difficulty" json:"difficulty"`
	Hidden                        bool                   `bson:"hidden" json:"hidden"`
	Hint                          string                 `bson:"hint" json:"hint"`
	Categories                    []string               `bson:"categories" json:"categories"`
	Descriptions                  map[string]Description `bson:"descriptions" json:"descriptions"`
	MemoryLimit                   int                    `bson:"memoryLimit" json:"memoryLimit"`
	ProblemCode                   interface{}            `bson:"problemCode" json:"problemCode"`
	ProblemID                     int                    `bson:"problemId" json:"problemId"`
	ProblemSetterUserID           int                    `bson:"problemSetterUserId" json:"problemSetterUserId"`
	SampleInput                   []string               `bson:"sampleInput" json:"sampleInput"`
	SampleOutput                  []string               `bson:"sampleOutput" json:"sampleOutput"`
	SampleTestCaseCount           int                    `bson:"sampleTestCaseCount" json:"sampleTestCaseCount"`
	TimeLimit                     int                    `bson:"timeLimit" json:"timeLimit"`
	Title                         string                 `bson:"title" json:"title"`
	Tried                         bool                   `json:"tried"`
	IsAccepted                    bool                   `json:"isAccepted"`
	AcceptedSubmission            int                    `json:"acceptedSubmission"`
	CompileErrorSubmission        int                    `json:"compileErrorSubmission"`
	MemoryLimitExceededSubmission int                    `json:"memoryLimitExceededSubmission"`
	OutputLimitExceededSubmission int                    `json:"outputLimitExceededSubmission"`
	RunTimeErrorSubmission        int                    `json:"runTimeErrorSubmission"`
	SubmissionErrorSubmission     int                    `json:"submissionErrorSubmission"`
	TimeLimitExceededSubmission   int                    `json:"timeLimitExceededSubmission"`
	TotalSubmission               int                    `json:"totalSubmission"`
	WrongAnswerSubmission         int                    `json:"wrongAnswerSubmission"`
}

type PStat struct {
	ID                   bson.ObjectId `bson:"_id,omitempty" json:"-"`
	AcceptedSubmission   int           `bson:"acceptedSubmission" json:"acceptedSubmission"`
	Difficulty           float64       `bson:"difficulty" json:"difficulty"`
	DistinctAcceptedUser int           `bson:"distinctAcceptedUser" json:"distinctAcceptedUser"`
	ProblemCode          interface{}   `bson:"problemCode" json:"problemCode"`
	ProblemID            int           `bson:"problemId" json:"problemId"`
	Title                string        `bson:"title" json:"title"`
	TotalSubmission      int           `bson:"totalSubmission" json:"totalSubmission"`
	LocaleAvailableEN    bool          `bson:"localeAvailableEN" json:"localeAvailableEN"`
	LocaleAvailableVI    bool          `bson:"localeAvailableVI" json:"localeAvailableVI"`
	Tried                bool          `json:"tried"`
	IsAccepted           bool          `json:"isAccepted"`
	IsBookmarked         bool          `json:"isBookmarked"`
}

type Description struct {
	Problem string `bson:"problem" json:"problem"`
	Input   string `bson:"input" json:"input"`
	Output  string `bson:"output" json:"output"`
	Note    string `bson:"note" json:"note"`
}

func PDetailList(query *bson.M, problems *[]PDetail) (err error) {
	sess := session.Copy()
	defer sess.Close()
	coll := sess.DB(conn.Database).C(conf.CDetail)

	return coll.Find(query).All(problems)
}

func PStatList(query *bson.M, stats *[]PStat) (err error) {
	sess := session.Copy()
	defer sess.Close()
	coll := sess.DB(conn.Database).C(conf.CStat)

	return coll.Find(query).All(stats)
}

func PDetailFind(query *bson.M, problem *PDetail) (err error) {
	sess := session.Copy()
	defer sess.Close()
	coll := sess.DB(conn.Database).C(conf.CDetail)

	return coll.Find(query).One(problem)
}
