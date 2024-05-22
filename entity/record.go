package entity

type Record struct {
	ID          int    `json:"id"`
	HeroName    string `json:"hero_name"`
	VillainName string `json:"villain_name"`
	Description string `json:"description"`
	EventTime   string `json:"event_time"`
}
