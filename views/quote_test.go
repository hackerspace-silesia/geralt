package views

import "testing"

func TestQuote(t *testing.T) {

	t.Run("Check if Quote has any characters", func(t *testing.T) {
		quote_handler := NewQuoteHandler(nil, "")
		quote := quote_handler.GetRandomQuote()
		if len(quote) <= 0 {
			t.Error("Invalid Quote")
		}

	})
}
