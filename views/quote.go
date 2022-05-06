package views

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackerspace-silesia/geralt/services"
	"github.com/slack-go/slack"
)

type QuoteHandler struct {
	apiClient      SlackClient
	secretVerifier SecretsVerifier
	quoteService   *services.QuoteService
}

func NewQuoteHandler(slackClient SlackClient, secrets SecretsVerifier) *QuoteHandler {

	return &QuoteHandler{
		apiClient:      slackClient,
		secretVerifier: secrets,
		quoteService:   services.NewQuoteService(),
	}
}

func (handler *QuoteHandler) QuoteServe(c *gin.Context) {
	if err := handler.secretVerifier.VerifySecret(c); err != nil {
		return
	}
	command, err := slack.SlashCommandParse(c.Copy().Request)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	switch command.Command {
	case "/quote":
		quote := handler.quoteService.GetRandomQuote()
		_, _, _, err := handler.apiClient.SendMessage(
			command.ChannelName,
			slack.MsgOptionText(quote, false),
		)
		if err != nil {
			fmt.Println(err)
		}
	}

}
