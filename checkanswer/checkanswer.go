package checkanswer

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/higashi000/oneprog-oneans/problem"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Answers struct {
	Id         string   `json:"id" bson:"id"`
	YourAnswer []string `json:"youranswer" bson:"youranswer"`
}

type Result struct {
	Text          string `json:"text"`
	YourAnswer    string `json:"youranswer"`
	CorrectAnswer string `json:"correctanswer"`
	Status        string `json:"Status"`
}

func CheckAnswer(r *gin.Engine, db *mongo.Database) {
	r.PUT("oneprog-oneans/checkanswer", func(c *gin.Context) {
		var answers Answers
		c.BindJSON(&answers)

		fmt.Println(answers)

		var setlist problem.SetList

		problemColle := db.Collection("problem")

		objectID, _ := primitive.ObjectIDFromHex(answers.Id)
		err := problemColle.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&setlist)
		if err != nil {
			log.Println("hoge", err)
		}

		var result []Result

		for i := 0; i < len(setlist.Problems); i++ {
			if setlist.Problems[i].CorrectAnswer == answers.YourAnswer[i] {
				result = append(result,
					Result{
						setlist.Problems[i].ProblemText,
						answers.YourAnswer[i],
						setlist.Problems[i].CorrectAnswer,
						"正解",
					})
			} else {
				result = append(result,
					Result{
						setlist.Problems[i].ProblemText,
						answers.YourAnswer[i],
						setlist.Problems[i].CorrectAnswer,
						"不正解",
					})
			}
		}

		c.JSON(http.StatusOK, result)
	})
}
