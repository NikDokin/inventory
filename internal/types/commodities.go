package types

type CommoditiesFilters struct {
	Name string `json:"name"`
}

type Commodity struct {
	Id          string  `json:"id"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Name        string  `json:"name"`
	PackageForm string  `json:"packageForm"`
	Price       float32 `json:"price"`
	Quantity    int     `json:"quantity"`
	Sku         string  `json:"sku"`
}
