package models

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
