package entity

type Villain struct {
	ID          int    `json:"id"`
	VillainName string `json:"villain_name"`
	Universe    string `json:"universe"`
	ImageURL    string `json:"image_url"`
}
