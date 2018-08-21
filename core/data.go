package core

type Event struct {
	StartDateTime string   `json:"start_date_time"`
	EndDateTime   string   `json:"end_date_time"`
	Location      string   `json:"location"`
	HTMLLink      string   `json:"html_link"`
	Summary       string   `json:"summary"`
	Title         string   `json:"title"`
	Tags          []string `json:"tags"`
}

type Events []Event

type ReturningData struct {
	Data interface{} `json:"data"`
}
