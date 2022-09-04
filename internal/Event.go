package internal

type Event struct {
	Name      string `json:"name"`
	Day       string `json:"day"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Week      string `json:"week"`
}
