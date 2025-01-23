package types

type CommoditiesFilters struct {
	Name string `json:"name"`
}

type Commodity struct {
	Id          string  `json:"id"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Name        string  `json:"name"`
	Package     string  `json:"Package"`
	Price       float32 `json:"price"`
	Quantity    int     `json:"quantity"`
	Sku         string  `json:"sku"`
}
