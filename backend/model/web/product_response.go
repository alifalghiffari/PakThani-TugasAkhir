package web

type ProductResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	CategoryId  int    `json:"categoryId"`
	Category    string `json:"category"`
	Quantity    int    `json:"Quantity"`
	Slug        string `json:"slug"`
}
