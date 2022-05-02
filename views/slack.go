package views

import "github.com/slack-go/slack"

type SlackClient interface {
	SendMessage(channel string, options ...slack.MsgOption) (string, string, string, error)
}
