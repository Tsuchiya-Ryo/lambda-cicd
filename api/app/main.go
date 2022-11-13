package main

import (
	"lambda-cicd/api/adaptor"
	"lambda-cicd/api/router"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("api/env/.env")
	if err != nil {
		log.Fatal(err)
	}
	e := gin.New()
	s3repo := adaptor.NewS3Repository()
	lambdaOp := adaptor.NewLambdaOperator()
	router.InitRouterA(e, *s3repo, *lambdaOp)
	e.Run(":8088")
}
