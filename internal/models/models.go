package models

type Product struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Price int    `json:"price,omitempty"`
}

type GetByName struct {
	SearchName string `json:"searchName"`
}
