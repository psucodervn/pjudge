package conf

const (
	JudgeHost = "http://192.168.99.100:8888"
)

const (
	CDetail     = "problem_detail"
	CStat       = "problem_stat"
	CSubmission = "submission"
	CUser       = "user"
	CIDs        = "ids"
)

const (
	ProblemPerPage  = 50
	ContestPerPage  = 100
	SolutionPerPage = 30
	UserPerPage     = 50
)

const (
	EN = "EN"
	VI = "VI"
)

const (
	Success = "success"
	Failed  = "failed"
	Warning = "warning"
)

const (
	JudgePD  = 0  //Pending
	JudgeRJ  = 1  //Running & judging
	JudgeCE  = 2  //Compile Error
	JudgeAC  = 3  //Accepted
	JudgeRE  = 4  //Runtime Error
	JudgeWA  = 5  //Wrong Answer
	JudgeTLE = 6  //Time Limit Exceeded
	JudgeMLE = 7  //Memory Limit Exceeded
	JudgeOLE = 8  //Output Limit Exceeded
	JudgePE  = 9  //Presentation Error
	JudgeNA  = 10 //System Error
	JudgeRPD = 11 //Rejudge Pending
)

const (
	LanguageNA   = 0 // None
	LanguageC    = 1 // GNU C
	LanguageCPP  = 2 // GNU C++
	LanguageJAVA = 3 // Java
)

var (
	Verdicts = []string{
		"Pending",
		"Running & judging",
		"Compile Error",
		"Accepted",
		"Runtime Error",
		"Wrong Answer",
		"Time Limit Exceeded",
		"Memory Limit Exceeded",
		"Output Limit Exceeded",
		"Presentation Error",
		"System Error",
		"Rejudge Pending",
	}
	Languages  = []string{"Text", "GNU C", "GNU C++", "Java"}
	LanguageID = map[string]int{
		"Text": 0, "GNU C": 1, "GNU C++": 2, "Java": 3,
	}
)

const (
	PrivilegeNA = 0 //None
	PrivilegePU = 1 //Primary User
	PrivilegeTC = 2 //Teacher
	PrivilegeAD = 3 //Admin
)
