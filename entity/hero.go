package entity

type Hero struct {
	ID       int    `json:"id"`
	HeroName string `json:"hero_name"`
	Universe string `json:"universe"`
	Skill    string `json:"skill"`
	ImageURL string `json:"image_url"`
}
