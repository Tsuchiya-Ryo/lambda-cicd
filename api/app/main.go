package main

import (
	"log"
	"os"

	"github.com/Tsuchiya-Ryo/lambda-cicd/api/adaptor"
	"github.com/Tsuchiya-Ryo/lambda-cicd/api/router"

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
		return godotenv.Load("env/prd.env")
	case "stg":
		return godotenv.Load("env/stg.env")
	default:
		return godotenv.Load("env/.env")
	}
}
