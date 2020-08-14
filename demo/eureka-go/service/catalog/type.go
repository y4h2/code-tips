package catalog

type catalogItem struct {
	ProductID       int    `json:"product_id"`
	SKU             string `json:"sku"`
	Description     string `json:"description"`
	Price           uint32 `json:"price"`
	ShipsWithin     int    `json:"ships_within"`
	QuantityInStock int    `json:"qty_in_stock"`
}

type fulfillmentStatus struct {
	SKU             string `json:"sku"`
	ShipsWithin     int    `json:"ships_within"`
	QuantityInStock int    `json:"qty_in_stock"`
}
