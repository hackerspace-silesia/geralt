package views

import (
	"strconv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"github.com/gin-gonic/gin"
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

type FinancesHandler struct {
	apiClient     *slack.Client
	signingSecret string
}

func NewFinancesHandler(slackClient *slack.Client, secret string) *FinancesHandler {
	fmt.Println("creating NewFinancesHandler")

	return &FinancesHandler{
		apiClient:     slackClient,
		signingSecret: secret,
	}
}

func (handler *FinancesHandler) FinancesServe(c *gin.Context) {
	if err := VerifySecret(c, handler.signingSecret); err != nil {
		return
	}
	command, err := slack.SlashCommandParse(c.Copy().Request)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if command.Command != "/finances" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	_, _, _, e := handler.apiClient.SendMessage(
		command.ChannelName,
		slack.MsgOptionText("1000 zl", false),
	)
	if e != nil {
		fmt.Println(e)
	}
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
