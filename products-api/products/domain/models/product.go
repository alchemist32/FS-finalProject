package models

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	BarCode     string `json:"barCode"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock,omitempty"`
}
