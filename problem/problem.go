package problem

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegistProblem(r *gin.Engine, db *mongo.Database) {
	r.POST("oneprog-oneans/registproblem", func(c *gin.Context) {
		var setlist SetList

		c.BindJSON(&setlist)

		problemColle := db.Collection("problem")

		res, err := problemColle.InsertOne(context.Background(), setlist)
		if err != nil {
			log.Println("insert collection:", err)
		}

		log.Println(res)

		c.JSON(http.StatusOK, setlist)
	})
}
