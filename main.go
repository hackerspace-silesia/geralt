package main

import (
	"strconv"
	"os"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/hackerspace-silesia/geralt/views"
	"github.com/slack-go/slack"
)

type Table struct {
	caption string
	values []TableEntry
}

type TableEntry struct {
	name string 
	value float64
}

func main() {
	tables := ParseFinancialTable("https://finanse.hs-silesia.pl")
	fmt.Println(len(tables))

	token := os.Getenv("SLACK_AUTH_TOKEN")
	signingSecret := os.Getenv("SLACK_SIGNING_SECRET")
	slackClient := slack.New(token)
	r := gin.Default()
	r.GET("/healthcheck", views.HealtcheckHandler)
	r.POST("/commands", views.NewQuoteHandler(slackClient, signingSecret).QuoteServe)
	r.Run()
}

func ParseFinancialTable(url string) []Table {
	document, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}
	raw_tables := document.Find("div.month table")
	all_tables := []Table{}
	raw_tables.Each(func(i int, s *goquery.Selection)  {
		caption := s.Find("caption").Text()
		table := Table{caption: caption, values: []TableEntry{}}
	    s.Find("tbody tr").Each(func(i int, s *goquery.Selection)  {
			name := s.Find("td .name").Text()
			value := StringToFloat(s.Find("td .value"))
			table.values = append(table.values, TableEntry{name: name, value: value})
		})
		all_tables = append(all_tables, table)
	})

	for i:=0; i<len(all_tables); i++ {
		fmt.Println(all_tables[i].caption)
	}
	return all_tables
}

func StringToFloat(s *goquery.Selection) float64 {
	value_string := s.Text()
	if value_string == "" {
		return 0
	}
	value_float, err := strconv.ParseFloat(value_string, 64)
	if err != nil {
		panic(err)
	}
	return value_float
}
