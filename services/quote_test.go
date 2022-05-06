package services

import "testing"

func TestQuote(t *testing.T) {

	t.Run("Check if Random quote has any characters", func(t *testing.T) {
		quote_handler := NewQuoteService()
		quote := quote_handler.GetRandomQuote()
		if len(quote) <= 0 {
			t.Error("Invalid Quote")
		}
	})
}
