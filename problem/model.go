package problem

type SetList struct {
	Title    string    `json:"title"`
	Problems []Problem `json:"problems"`
}

type Problem struct {
	ProblemText   string `json:"problemtext"`
	CorrectAnswer string `json:"correctanswer"`
}
