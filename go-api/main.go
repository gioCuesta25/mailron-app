package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type ZincSearchResponse struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			ID     string `json:"_id"`
			Source struct {
				Content string `json:"content"`
				From    string `json:"from"`
				Subject string `json:"subject"`
				To      string `json:"to"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type Response struct {
	Total int        `json:"total"`
	Items []MailItem `json:"items"`
}

type MailItem struct {
	Id      string `json:"id"`
	Content string `json:"content"`
	From    string `json:"from"`
	Subject string `json:"subject"`
	To      string `json:"to"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/documents/match", MatchPhrase)

	http.ListenAndServe(":4000", r)
}

func MatchPhrase(w http.ResponseWriter, r *http.Request) {

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
	term := r.URL.Query().Get("term")
	from := r.URL.Query().Get("from")

	if term != "" {
		searchType = "matchphrase"
	}

	reqExpression := fmt.Sprintf(query, searchType, term, from)

	req, err := http.NewRequest("POST", "http://localhost:4080/api/test/_search", strings.NewReader(reqExpression))

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

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var result ZincSearchResponse

	err = json.Unmarshal(body, &result)

	if err != nil {
		log.Fatal(err)
	}

	var response Response

	response.Total = result.Hits.Total.Value

	for _, item := range result.Hits.Hits {
		var newItem MailItem
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

	w.Write(data)

}
