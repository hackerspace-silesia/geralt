package views

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
	"github.com/stretchr/testify/mock"
)

func TestQuote(t *testing.T) {

	slashCommand := &SlashCommand{
		Token:          "",
		TeamID:         "",
		TeamDomain:     "",
		EnterpriseID:   "",
		EnterpriseName: "",
		ChannelID:      "AZ12345S",
		ChannelName:    "bot-testing",
		UserID:         "",
		UserName:       "",
		Command:        "/quote",
		Text:           "",
		ResponseURL:    "",
		TriggerID:      "",
		APIAppID:       "",
	}
	form := url.Values{}
	encoder := schema.NewEncoder()
	if err := encoder.Encode(slashCommand, form); err != nil {
		t.Error("Cannot encode PostForm command")
	}
	t.Run("Check Handler is sending a message", func(t *testing.T) {
		slackClientMock := new(SlackClientMock)
		slackClientMock.On("SendMessage", slashCommand.ChannelName, mock.Anything).Return("", "", "", nil)

		verifierMock := new(SecretsVerifierMock)
		verifierMock.On("VerifySecret", mock.Anything).Return(nil)

		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		ctx.Request, _ = http.NewRequest("POST", "/commmand", new(bytes.Buffer))
		ctx.Request.PostForm = form
		quote_handler := NewQuoteHandler(slackClientMock, verifierMock)
		quote_handler.QuoteServe(ctx)

		slackClientMock.AssertNumberOfCalls(t, "SendMessage", 1)
	})
}
