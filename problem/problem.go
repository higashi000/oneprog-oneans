package problem

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetSetList(r *gin.Engine, db *mongo.Database) {
	r.GET("oneprog-oneans/getsetlist", func(c *gin.Context) {
		var setlists []TitleID

		problemColle := db.Collection("problem")

		findOptions := options.Find()
		findOptions.SetLimit(5)

		cur, err := problemColle.Find(context.TODO(), bson.D{{}}, findOptions)
		if err != nil {
			log.Fatal(err)
		}

		for cur.Next(context.Background()) {
			var tmp TitleID
			cur.Decode(&tmp)

			setlists = append(setlists, tmp)
		}

		log.Println(setlists)
		c.JSON(http.StatusOK, setlists)
	})
}

func ChallengeProblem(r *gin.Engine, db *mongo.Database) {
	r.GET("oneprog-oneans/challenge/:id/:nowProblemNum", func(c *gin.Context) {
		problemColle := db.Collection("problem")

		id := c.Param("id")
		nowNum := c.Param("nowProblemNum")

		objectID, _ := primitive.ObjectIDFromHex(id)

		var setlist SetList

		err := problemColle.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&setlist)
		if err != nil {
			log.Println(err)
		}

		problemNum := len(setlist.Problems)
		num, err := strconv.Atoi(nowNum)
		if err != nil {
			log.Println(err)
		}

		type Status struct {
			Status string `json":status"`
		}

		if num >= problemNum {
			c.JSON(http.StatusOK, Status{"finish"})
		} else {
			c.JSON(http.StatusOK, setlist.Problems[num])
		}
	})
}
