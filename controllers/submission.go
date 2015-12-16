package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/psufighter/pjudge/conf"
	"github.com/psufighter/pjudge/models"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type SubmissionController struct {
	beego.Controller
}

func (c *SubmissionController) URLMapping() {
	c.Mapping("ViewLanguage", c.ViewLanguage)
}

// @router /submission/view/language [post]
func (c *SubmissionController) ViewLanguage() {
	c.Data["json"] = conf.Languages
	c.ServeJson()
}

// @router /submission/submit/:pid/text [post]
func (c *SubmissionController) SubmitText() {
	sessionUser := c.GetSession("user")
	if sessionUser == nil {
		makeJSONMessage(&c.Controller, "Must be logged in to submit!")
		return
	}
	user := sessionUser.(models.User)

	language := c.Input().Get("language")
	code := c.Input().Get("solution")
	pid := c.Ctx.Input.Param(":pid")
	problemID, _ := strconv.Atoi(pid)

	var prob models.PDetail
	if err := models.PDetailFind(&bson.M{"problemId": problemID}, &prob); err != nil {
		makeJSONMessage(&c.Controller, "Problem ID "+pid+" not found!")
		return
	}

	var subID int
	models.LastID(conf.CSubmission, &subID)
	sub := models.Submission{}
	sub.Language = language
	sub.LanguageID, _ = conf.LanguageID[language]
	sub.Code = code
	sub.ProblemID = prob.ProblemID
	sub.ProblemName = prob.Title
	// TODO: add problem name
	sub.SubmissionID = subID + 1
	sub.SubmitTime = time.Now().UnixNano() / 1000000
	sub.VerdictID = conf.JudgePD
	sub.Verdict = conf.Verdicts[sub.VerdictID]
	sub.Username = user.Username
	sub.UserID = user.UserID

	if err := models.SubmissionInsert(&sub); err == nil {
		models.IncreaseID(conf.CSubmission)
		makeJSONMessage(&c.Controller, "Success")
	} else {
		makeJSONMessage(&c.Controller, err.Error())
		return
	}

	// post to judger
	go func() {
		info := make(map[string]interface{})
		info["Sid"] = sub.SubmissionID
		info["Pid"] = prob.ProblemID
		info["OJ"] = ""
		info["Rejudge"] = false
		data, _ := json.Marshal(info)
		_, err := http.Post(conf.JudgeHost, "application/json", strings.NewReader(string(data)))
		if err != nil {
			beego.Error("Post to judger: ", err.Error())
		}
	}()
}

// @router /submission/view [post]
func (c *SubmissionController) View() {
	sort := "-submitTime"
	limit := conf.SolutionPerPage
	var subs []models.Submission
	if err := models.SubmissionList(nil, &subs, nil, &sort, nil, &limit); err == nil {
		c.Data["json"] = subs
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @router /submission/view/:sid [post]
func (c *SubmissionController) ViewSubmission() {
	sessionUser := c.GetSession("user")
	if sessionUser == nil {
		makeJSONMessage(&c.Controller, "Must be logged in to submit!")
		return
	}
	user := sessionUser.(models.User)

	sid := c.Ctx.Input.Param(":sid")
	submissionID, _ := strconv.Atoi(sid)
	var sub models.Submission
	models.SubmissionFind(&bson.M{"submissionId": submissionID}, &sub)

	beego.Error(user.UserID, sub.UserID, submissionID)

	if user.UserID != sub.UserID {
		makeJSONMessage(&c.Controller, "Cannot view submission of other user!")
		return
	}

	sub.Verdict = conf.Verdicts[sub.VerdictID]
	makeJSONMessage(&c.Controller, sub)
}

// @router /submission/view/user/count/:uid [post]
func (c *SubmissionController) UserCount() {
	uid := c.Ctx.Input.Param(":uid")
	userID, _ := strconv.Atoi(uid)

	count, _ := models.Count(conf.CSubmission, &bson.M{"userId": userID})

	makeJSONMessage(&c.Controller, count)
}

// @router /submission/view/user/:p/:uid [post]
func (c *SubmissionController) UserView() {
	uid := c.Ctx.Input.Param(":uid")
	userID, _ := strconv.Atoi(uid)
	p := c.Ctx.Input.Param(":p")
	page, _ := strconv.Atoi(p)

	sort := "-_id"
	skip := (page - 1) * 5
	limit := 5
	var subs []models.Submission
	models.SubmissionList(&bson.M{"userId": userID}, &subs, nil, &sort, &skip, &limit)

	makeJSONMessage(&c.Controller, subs)
}
