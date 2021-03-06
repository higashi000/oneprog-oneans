package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/higashi000/oneprog-oneans/checkanswer"
	"github.com/higashi000/oneprog-oneans/problem"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		log.Println(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Println(err)
	}

	db := client.Database("oneprog-oneans")

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	problem.RegistProblem(r, db)
	problem.GetSetList(r, db)
	problem.ChallengeProblem(r, db)
	checkanswer.CheckAnswer(r, db)

	r.Run()
}
