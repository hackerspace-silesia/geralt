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
	slackClient := slack.New(token, slack.OptionDebug(true))
	r := gin.Default()
	r.GET("/healthcheck", views.HealtcheckHandler)
	r.POST("/commands", views.NewQuoteHandler(slackClient, signingSecret).QuoteServe)
	r.POST("/finances", views.NewFinancesHandler(slackClient, signingSecret).FinancesServe)
	r.Run()
}
