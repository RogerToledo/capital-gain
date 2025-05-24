package service

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/me/capital-gain/entity"
)

func ProcessTransactions(transactions []entity.Transaction) string {
	var (
		taxs                   []entity.Out
		tax 				   float64
		profit                 float64
		currentStocks          float64
		currentWeightedAverage float64
	)

	for _, in := range transactions {
		if in.Operation == "buy" {
			if currentWeightedAverage == 0 {
				currentWeightedAverage = in.UnitCost
			}

			profit += calculateProfit(currentWeightedAverage, in.UnitCost, in.Quantity)

			currentWeightedAverage = calcWeightedAverage(currentStocks, in.Quantity, currentWeightedAverage, in.UnitCost)

			taxs = append(taxs, entity.Out{Tax: 0.00})

			currentStocks += in.Quantity
		}

		if in.Operation == "sell" {
			newProfit := calculateProfit(currentWeightedAverage, in.UnitCost, in.Quantity)

			tax, profit = calculateTax(currentWeightedAverage, profit, newProfit, in)

			taxs = append(taxs, entity.Out{Tax: tax})

			currentStocks -= in.Quantity
		}
	}

	jsonData, err := json.Marshal(taxs)
	if err != nil {
		return err.Error()
	}

	return string(jsonData)
}

func calculateTax(currentWeightedAverage, profit, newProfit float64, in entity.Transaction) (float64, float64) {
	var (
		taxRate = 0.20
		tax     float64
	)

	if in.UnitCost > currentWeightedAverage {
		if profit > 0 {
			operationCost := in.UnitCost * in.Quantity
			
			if operationCost > 20000.00 {
				profit += newProfit

				tax = profit * taxRate

				return tax, profit
			} else {
				return 0.00, profit
			}
		} else {
			profit += newProfit

			return 0.00, profit
		}
	}

	if in.UnitCost < currentWeightedAverage {
		profit += newProfit

		return 0.00, profit
	}

	return 0.00, profit
}

func calculateProfit(currentWeightedAverage, unitCost, quantity float64) float64 {
	return (unitCost - currentWeightedAverage) * quantity
}

func calcWeightedAverage(currentStocks, stocksBought, currentWeightedAverage, stocksCost float64) float64 {
	one := currentStocks * currentWeightedAverage
	two := stocksBought * stocksCost
	three := currentStocks + stocksBought

	fmt.Printf("((%f * %f) + (%f * %f)) / (%f + %f)\n", currentStocks, currentWeightedAverage, stocksBought, stocksCost, currentStocks, stocksBought)
	fmt.Printf("((%f) + (%f)) / %f\n", one, two, three)
	four := one + two
	fmt.Printf("(%f) / %f\n", four, three)
	wa := ((currentStocks * currentWeightedAverage) + (stocksBought * stocksCost)) / (currentStocks + stocksBought)

	return math.Round(wa * 100) / 100
}
