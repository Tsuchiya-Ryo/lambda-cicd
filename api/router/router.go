package router

import (
	"net/http"
	"os"

	"github.com/Tsuchiya-Ryo/lambda-cicd/api/adaptor"
	"github.com/Tsuchiya-Ryo/lambda-cicd/api/logic"
	"github.com/Tsuchiya-Ryo/lambda-cicd/api/model"

	"github.com/gin-gonic/gin"
)

type Router struct {
	S3Repository   adaptor.S3Repository
	LambdaOperator adaptor.LambdaOperator
}

func InitRouter(e *gin.Engine, s3repo adaptor.S3Repository, lambaOp adaptor.LambdaOperator) {
	router := &Router{
		S3Repository:   s3repo,
		LambdaOperator: lambaOp,
	}
	e.GET("/invoke_lambda", router.Invoke)
	e.GET("/health_check", router.HealthCheck)
}

func (r *Router) Invoke(c *gin.Context) {
	from := c.Request.URL.Query().Get("from_date")
	to := c.Request.URL.Query().Get("to_date")
	method := c.Request.URL.Query().Get("method")

	models, err := logic.GenerateData(from, to)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}
	body, err := logic.ToBody(models)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}
	err = r.S3Repository.PutObject(os.Getenv("BUCKET_NAME"), os.Getenv("S3_FILE_NAME"), body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}
	resp, err := r.LambdaOperator.Invoke(model.InvokeInput{Key: os.Getenv("S3_FILE_NAME"), Method: method})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "accepted", "lambda_response": string(resp.Payload)})
}

func (r *Router) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "PogU")
}
