package services

import "testing"

func TestEventService(t *testing.T) {
	t.Run("There are upcoming events", func(t *testing.T) {
		eventService := NewEventService("https://hs-silesia.pl/feeds/atom.xml")
		_, err := eventService.GetUpcomingEvents()
		if err != nil {
			t.Error(err)
			return
		}
	})
	t.Run("There are any events", func(t *testing.T) {
		eventService := NewEventService("https://hs-silesia.pl/feeds/atom.xml")
		events, err := eventService.GetLatestEvents()
		if err != nil {
			t.Error(err)
			return
		}
		if len(events) == 0 {
			t.Error("There is expected at least one event!")
			return
		}
	})
}
