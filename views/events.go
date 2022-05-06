package views

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/hackerspace-silesia/geralt/services"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

type SlackEventsHandler struct {
	apiClient      SlackClient
	secretVerifier SecretsVerifier
	eventService   services.EventsService
}

func NewSlackEventsHandler(slackClient SlackClient, secrets SecretsVerifier) *SlackEventsHandler {
	return &SlackEventsHandler{
		apiClient:      slackClient,
		secretVerifier: secrets,
		eventService:   *services.NewEventService("http://hs-silesia.pl/feeds/atom.xml"),
	}
}

func (handler *SlackEventsHandler) EventServe(c *gin.Context) {
	if err := handler.secretVerifier.VerifySecret(c); err != nil {
		return
	}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Errorf("Problem with accessing body!")
	}
	eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
	if err != nil {
		fmt.Print(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		event := eventsAPIEvent.Data.(*slackevents.EventsAPIURLVerificationEvent)
		c.Header("Content-Type", "text")
		c.String(http.StatusOK, event.Challenge)
	}

	if eventsAPIEvent.InnerEvent.Type == slackevents.AppMention {
		event := eventsAPIEvent.InnerEvent.Data.(*slackevents.AppMentionEvent)
		client := handler.apiClient.(*slack.Client)
		client.PostMessage("")
		re := regexp.MustCompile(`^<@U036Q3LGRMH> (?P<command>[a-z_]+)`)
		commandGroupIndex := re.SubexpIndex("command")
		matches := re.FindStringSubmatch(event.Text)
		fmt.Println(event.Text)
		if len(matches) > 1 {
			switch matches[commandGroupIndex] {
			case "upcoming_events":
				handler.ShowUpcomingEvents(event.Channel)
			case "all_events":
				handler.ShowAllEvents(event.Channel)
			default:
				fmt.Println("Command not supported!")
			}
		} else {
			fmt.Println("Couldn't parse command!")
		}

	}
}

func (handler *SlackEventsHandler) ShowUpcomingEvents(channel string) {
	events, err := handler.eventService.GetUpcomingEvents()
	if err != nil {
		fmt.Errorf("Something went wrong when fetching events!")
		return
	}
	eventBlocks := []*slack.TextBlockObject{}
	for _, event := range events {
		message := fmt.Sprintf("%s - %s", event.Title, event.Date.Format("2006-01-02"))
		eventBlocks = append(eventBlocks, slack.NewTextBlockObject("mrkdwn", message, false, false))
	}
	headerText := slack.NewTextBlockObject("mrkdwn", "*Upcoming Hackerspace Events*", false, false)
	headerSection := slack.NewSectionBlock(headerText, eventBlocks, nil)
	handler.apiClient.SendMessage(channel, slack.MsgOptionBlocks(headerSection))
}

func (handler *SlackEventsHandler) ShowAllEvents(channel string) {
	events, err := handler.eventService.GetLatestEvents()
	if err != nil {
		fmt.Errorf("Something went wrong when fetching events!")
		return
	}
	eventBlocks := []*slack.TextBlockObject{}
	for _, event := range events {
		message := fmt.Sprintf("%s - %s", event.Title, event.Date.Format("2006-01-02"))
		eventBlocks = append(eventBlocks, slack.NewTextBlockObject("mrkdwn", message, false, false))
	}
	headerText := slack.NewTextBlockObject("mrkdwn", "*Upcoming Hackerspace Events*", false, false)
	headerSection := slack.NewSectionBlock(headerText, eventBlocks, nil)
	handler.apiClient.SendMessage(channel, slack.MsgOptionBlocks(headerSection))
}
