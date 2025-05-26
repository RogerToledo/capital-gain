package entity

type Transaction struct {
	Operation string  `json:"operation"`
	UnitCost  float64 `json:"unit-cost"`
	Quantity  float64 `json:"quantity"`
}

type Out struct {
	Tax float64 `json:"tax"`
}
