package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hackerspace-silesia/geralt/views"
	"github.com/slack-go/slack"
)

func main() {
	token := os.Getenv("SLACK_AUTH_TOKEN")
	signingSecret := os.Getenv("SLACK_SIGNING_SECRET")
	slackClient := slack.New(token)
	secretsVerifier := views.NewSlackSecretsVerifier(signingSecret)
	r := gin.Default()
	r.GET("/healthcheck", views.HealtcheckHandler)
	r.POST("/commands", views.NewQuoteHandler(slackClient, secretsVerifier).QuoteServe)
	//r.POST("/events", views.NewEventsHandler(slackClient, secretsVerifier).QuoteServe)
	r.Run()
}
