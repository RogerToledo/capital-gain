package service

import (
	"encoding/json"
	"math"

	"github.com/me/capital-gain/entity"
)

func ProcessTransactions(transactions [][]entity.Transaction) string {
	var (
		taxs                   []entity.Out
		output                 [][]entity.Out
		tax                    float64
		profit                 float64
		currentStocks          float64
		currentWeightedAverage float64
	)

	for _, transaction := range transactions {
		for _, in := range transaction {
			if in.Operation == "buy" {
				if currentWeightedAverage == 0 {
					currentWeightedAverage = in.UnitCost
				}

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

		output = append(output, taxs)
		taxs = nil
		profit = 0
	}

	jsonData, err := json.Marshal(output)
	if err != nil {
		return err.Error()
	}

	taxs = nil
	currentStocks = 0
	currentWeightedAverage = 0
	profit = 0
	tax = 0.00

	return string(jsonData)
}

func calculateTax(currentWeightedAverage, profit, newProfit float64, in entity.Transaction) (float64, float64) {
	var (
		taxRate = 0.20
		tax     float64
	)

	if in.UnitCost > currentWeightedAverage {
		if newProfit >= 0 {
			operationCost := in.UnitCost * in.Quantity

			if operationCost > 20000.00 {
				profit += newProfit

				if profit > 0 {
					tax = profit * taxRate
					profit = 0
				}

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
	profit := (unitCost - currentWeightedAverage) * quantity

	return profit
}

func calcWeightedAverage(currentStocks, stocksBought, currentWeightedAverage, stocksCost float64) float64 {
	wa := ((currentStocks * currentWeightedAverage) + (stocksBought * stocksCost)) / (currentStocks + stocksBought)

	return math.Round(wa*100) / 100
}
