package main

import (
	"context"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/danteay/test-slsgo/utils"
)

type Response struct {
	Data string `json:"data"`
}

func Handler(_ context.Context) (utils.Response, error) {
	body := Response{Data: os.Getenv("APP_NAME")}

	res, err := utils.NewResponse(http.StatusOK, body)
	if err != nil {
		return utils.Response{StatusCode: http.StatusInternalServerError}, err
	}

	//res.WithCors()
	return *res, nil
}

func main() {
	lambda.Start(Handler)
}
