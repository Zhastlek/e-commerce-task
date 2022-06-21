package models

type Product struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Price int    `json:"price,omitempty"`
}

// type Data struct {
// 	Status     string `json:"status"`
// 	SearchName string `json:"searchName"`
// 	Error      string `json:"error"`
// 	ProductId  int    `json:"productId"`
// }

type GetByName struct {
	SearchName string `json:"searchName"`
}

type GiveProduct struct {
	Product
}

type Status struct {
	Status    string `json:"status"`
	Error     string `json:"error,omitempty"`
	ProductId int    `json:"productId,omitempty"`
}
