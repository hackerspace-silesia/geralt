package services

import "math/rand"

type QuoteService struct {
	quotes []string
}

func NewQuoteService() *QuoteService {
	return &QuoteService{
		quotes: []string{
			"Why men throw their lives away attacking an armed witcher... I'll never know. Something about my face?",
			"If I have to choose between one evil or another, I'd rather not choose at all",
			"People,' Geralt turned his head, 'like to invent monsters and monstrosities. Then they seem less monstrous themselves."},
	}
}

func (handler *QuoteService) GetRandomQuote() string {
	n := rand.Intn(len(handler.quotes))
	quote := handler.quotes[n]
	return quote
}
