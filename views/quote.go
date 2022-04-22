package views

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
)

type QuoteHandler struct {
	quotes        []string
	apiClient     *slack.Client
	signingSecret string
}

func NewQuoteHandler(slackClient *slack.Client, secret string) *QuoteHandler {

	return &QuoteHandler{
		apiClient:     slackClient,
		signingSecret: secret,
		quotes: []string{
			"Why men throw their lives away attacking an armed witcher... I'll never know. Something about my face?",
			"If I have to choose between one evil or another, I'd rather not choose at all",
			"People,' Geralt turned his head, 'like to invent monsters and monstrosities. Then they seem less monstrous themselves."},
	}
}

func (handler *QuoteHandler) QuoteServe(c *gin.Context) {
	if err := VerifySecret(c, handler.signingSecret); err != nil {
		return
	}
	command, err := slack.SlashCommandParse(c.Copy().Request)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Println(command)
	switch command.Command {
	case "/quote":
		quote := handler.GetRandomQuote()
		_, _, _, err := handler.apiClient.SendMessage(
			command.ChannelName,
			slack.MsgOptionText(quote, false),
		)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (handler *QuoteHandler) GetRandomQuote() string {
	n := rand.Intn(len(handler.quotes))
	quote := handler.quotes[n]
	return quote
}
