package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	SLACK_API_TOKEN := os.Getenv("SLACK_API_TOKEN")
	SLACK_VERIFICATION_TOKEN := os.Getenv("SLACK_VERIFICATION_TOKEN")

	api := slack.New(SLACK_API_TOKEN)
	parsedEvent, err := slackevents.ParseEvent(json.RawMessage(request.Body), slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: SLACK_VERIFICATION_TOKEN}))
	if err != nil {
		log.Printf("Error parsing event: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	if parsedEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(request.Body), &r)
		if err != nil {
			log.Printf("Error unmarshalling response: %v", err)
			return events.APIGatewayProxyResponse{StatusCode: 500}, nil
		}
		return events.APIGatewayProxyResponse{Body: r.Challenge, StatusCode: 200}, nil
	}

	if parsedEvent.Type == slackevents.CallbackEvent {
		innerEvent := parsedEvent.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			_, _, err := api.PostMessage(ev.Channel, slack.MsgOptionText("こんにちは！どのようにお手伝いできますか？", false), slack.MsgOptionTS(ev.TimeStamp))
			if err != nil {
				log.Printf("Error posting message: %v", err)
				return events.APIGatewayProxyResponse{StatusCode: 500}, nil
			}
		}
	}

	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
