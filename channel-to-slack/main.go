package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"channel-to-slack/domain"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	incomingWebHookURL := os.Getenv("INCOMING_WEBHOOK_URL")

	channel, err := domain.NewChannel(request.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	slack := channel.ToSlack()

	slackParams, err := json.Marshal(slack)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	resp, _ := http.PostForm(
		incomingWebHookURL,
		url.Values{"payload": {string(slackParams)}},
	)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	println(string(body))
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}
