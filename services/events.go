package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
)

type Event struct {
	Title string
	Date  time.Time
}

type EventsService struct {
	atomUrl string
}

func NewEventService(url string) *EventsService {
	return &EventsService{
		atomUrl: url,
	}
}

func (e *EventsService) GetLatestEvents() ([]Event, error) {
	events := []Event{}
	parser := gofeed.NewParser()
	feed, err := parser.ParseURL(e.atomUrl)
	if err != nil {
		fmt.Errorf("Error: %s", err)
		return events, err
	}
	for _, item := range feed.Items {
		if strings.HasPrefix(item.Link, "/events/") {
			events = append(events, Event{
				Title: item.Title,
				Date:  *item.PublishedParsed,
			})
		}
	}
	return events, err
}

func (e *EventsService) GetUpcomingEvents() ([]Event, error) {
	upcomingEvents := []Event{}
	events, err := e.GetLatestEvents()
	if err != nil {
		return events, err
	}
	for _, event := range events {
		if time.Now().Before(event.Date) {
			upcomingEvents = append(upcomingEvents, event)
		}
	}

	return upcomingEvents, nil

}
