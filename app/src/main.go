package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	SLACK_API_TOKEN := os.Getenv("SLACK_API_TOKEN")
	SLACK_VERIFICATION_TOKEN := os.Getenv("SLACK_VERIFICATION_TOKEN")

	api := slack.New(SLACK_API_TOKEN)

	// Get the bot's user ID
	authTestResponse, err := api.AuthTest()
	if err != nil {
		log.Printf("Error getting bot user ID: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}
	botUserID := authTestResponse.UserID

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
			if ev.User == botUserID {
				return events.APIGatewayProxyResponse{StatusCode: 200}, nil
			}
			question := strings.TrimPrefix(ev.Text, "@AIニキ")
			responseText := GetAIMessage(question)
			_, _, err := api.PostMessage(ev.Channel, slack.MsgOptionText(responseText, false), slack.MsgOptionTS(ev.TimeStamp))
			if err != nil {
				log.Printf("Error posting message: %v", err)
				return events.APIGatewayProxyResponse{StatusCode: 500}, nil
			}
		}
	}

	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func GetAIMessage(question string) string {
	OPENAI_API_KEY := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(OPENAI_API_KEY)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}

	return resp.Choices[0].Message.Content
}

func main() {
	lambda.Start(HandleRequest)
}
