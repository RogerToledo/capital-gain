package entity

type Transaction struct {
	Operation string
	UnitCost  float64
	Quantity  float64
}

type Out struct {
	Tax float64 `json:"tax"`
}
