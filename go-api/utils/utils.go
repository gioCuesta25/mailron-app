package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Giovanni-Romana-Cuesta/go-api/models"
)

func MakeSearchRequest(term string, from string) *http.Response {
	query := `{
		"search_type": "%s",
		"query":
		{
			"term": "%s"
		},
		"from": %s,
		"max_results": 20,
		"_source": []
	}`

	searchType := "alldocuments"

	if term != "" {
		searchType = "matchphrase"
	}

	reqExpression := fmt.Sprintf(query, searchType, term, from)

	req, err := http.NewRequest("POST", "http://localhost:4080/api/enron_mails/_search", strings.NewReader(reqExpression))

	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth("admin", "Complex#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func ParseResponse(body []byte) []byte {
	var result models.ZincSearchResponse

	err := json.Unmarshal(body, &result)

	if err != nil {
		log.Fatal(err)
	}

	var response models.Response

	response.Total = result.Hits.Total.Value

	for _, item := range result.Hits.Hits {
		var newItem models.MailItem
		newItem.Content = item.Source.Content
		newItem.From = item.Source.From
		newItem.To = item.Source.To
		newItem.Subject = item.Source.Subject
		newItem.Id = item.ID

		response.Items = append(response.Items, newItem)
	}

	data, err := json.Marshal(response)

	if err != nil {
		log.Fatal(err)
	}

	return data
}
