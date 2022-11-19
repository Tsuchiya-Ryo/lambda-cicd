package main

import (
	"lambda-cicd/api/adaptor"
	"lambda-cicd/api/router"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := load_config()
	if err != nil {
		log.Fatal(err)
	}

	e := gin.New()
	s3repo := adaptor.NewS3Repository()
	lambdaOp := adaptor.NewLambdaOperator()
	router.InitRouterA(e, *s3repo, *lambdaOp)

	e.Run(":8088")
}

func load_config() error {
	env := os.Getenv("ENV")
	switch env {
	case "prd":
		return godotenv.Load("api/env/prd.env")
	case "stg":
		return godotenv.Load("api/env/stg.env")
	default:
		return godotenv.Load("api/env/.env")
	}
}
