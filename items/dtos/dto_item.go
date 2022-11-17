package dtos

type ItemDto struct {
	Tittle      string  `json:"tittle"`
	Id          string  `json:"id"`
	Seller      string  `json:"seller"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
	Pictures    string  `json:"pictures"`
	Description string  `json:"description"`
	State       string  `json:"state"`
	City        string  `json:"city"`
	Street      string  `json:"street"`
	Number      int     `json:"number"`
}

type ItemsDto []ItemDto
