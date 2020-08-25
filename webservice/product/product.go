package product

//Product information
type Product struct {
	ProductID      int    `json:"productid"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"priceperunit"`
	QuantityOnHand int    `json:"quantityonhand"`
	ProductName    string `json:"productname"`
}
