package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler Handle lambda
func Handler(_ context.Context) (events.APIGatewayProxyResponse, error) {
	// body := map[string]string{"data": os.Getenv("APP_NAME")}

	// res, err := utils.NewResponse(http.StatusOK, body)
	// if err != nil {
	// 	return utils.Response{StatusCode: http.StatusInternalServerError}, err
	// }

	// //res.WithCors()
	// return *res, nil

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       `{"ping": "pong"}`,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
