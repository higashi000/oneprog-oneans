package problem

type SetList struct {
	ID       string    `json:"_id" bson:"_id"`
	Title    string    `json:"title" bson:"title"`
	Problems []Problem `json:"problems" bson:"problems"`
}

type Problem struct {
	ProblemText   string `json:"problemtext" bson:"problemtext"`
	CorrectAnswer string `json:"correctanswer" bson:"correctanswer"`
}
