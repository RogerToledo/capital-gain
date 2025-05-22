package main

type In []struct {
	Operation string  `json:"operation"`
	UnitCost  float32 `json:"unit_cost"`
	Quantity  int     `json:"quantity"`
}

type Out []struct {
	Tax float32 `json:"tax"`
}
