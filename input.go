package main

import "github.com/me/capital-gain/entity"

var JasonIn = []entity.Transaction{
	{
		Operation: "buy",
		UnitCost:  10.0,
		Quantity:  10000,
	},
	{
		Operation: "sell",
		UnitCost:  20.0,
		Quantity:  5000,
	},
	{
		Operation: "buy",
		UnitCost:  20.0,
		Quantity:  10000,
	},
	{
		Operation: "sell",
		UnitCost:  10.0,
		Quantity:  5000,
	},
}
