package service

import (
	"testing"

	"github.com/me/capital-gain/entity"
)

func TestProcessTransactions(t *testing.T) {
	cases := []struct {
		name     string
		input    []entity.Transaction
		input1   []entity.Transaction
		expected string
	}{
		// {
		// 	name: "Case 1",
		// 	input: []entity.Transaction{
		// 		{
		// 			Operation: "buy",
		// 			UnitCost:  10.00,
		// 			Quantity:  100,
		// 		},
		// 		{
		// 			Operation: "sell",
		// 			UnitCost:  15.00,
		// 			Quantity:  50,
		// 		},
		// 		{
		// 			Operation: "sell",
		// 			UnitCost:  15.00,
		// 			Quantity:  50,
		// 		},
		// 	},
		// 	expected: `[{"tax":0},{"tax":0},{"tax":0}]`,
		// },
		{
			name: "Case 2",
			input: []entity.Transaction{
				{
					Operation: "buy",
					UnitCost:  10.00,
					Quantity:  10000,
				},
				{
					Operation: "sell",
					UnitCost:  20.00,
					Quantity:  5000,
				},
				{
					Operation: "sell",
					UnitCost:  5.00,
					Quantity:  5000,
				},
			},
			expected: `[{"tax":0},{"tax":10000},{"tax":0}]`,
		},
		//Must to do Case 1 + Case 2
		// {
		// 	name: "Case 3",
		// 	input: []entity.Transaction{
		// 		{
		// 			Operation: "buy",
		// 			UnitCost:  10.00,
		// 			Quantity:  10000,
		// 		},
		// 		{
		// 			Operation: "sell",
		// 			UnitCost:  5.00,
		// 			Quantity:  5000,
		// 		},
		// 		{
		// 			Operation: "sell",
		// 			UnitCost:  20.00,
		// 			Quantity:  3000,
		// 		},
		// 	},
		// 	expected: `[{"tax":0},{"tax":0},{"tax":1000}]`,
		// },
		// {
		// 	name: "Case 4",
		// 	input: []entity.Transaction{
		// 		{
		// 			Operation: "buy",
		// 			UnitCost:  10.00,
		// 			Quantity:  10000,
		// 		},
		// 		{
		// 			Operation: "buy",
		// 			UnitCost:  25.00,
		// 			Quantity:  5000,
		// 		},
		// 		{
		// 			Operation: "sell",
		// 			UnitCost:  15.00,
		// 			Quantity:  10000,
		// 		},
		// 	},
		// 	expected: `[{"tax":0},{"tax":0},{"tax":0}]`,
		// },
		// {
		// 	name: "Case 5",
		// 	input: []entity.Transaction{
		// 		{
		// 			Operation: "buy",
		// 			UnitCost:  10.00,
		// 			Quantity:  10000,
		// 		},
		// 		{
		// 			Operation: "buy",
		// 			UnitCost:  25.00,
		// 			Quantity:  5000,
		// 		},
		// 		{
		// 			Operation: "sell",
		// 			UnitCost:  15.00,
		// 			Quantity:  10000,
		// 		},
		// 		{
		// 			Operation: "sell",
		// 			UnitCost:  25.00,
		// 			Quantity:  5000,
		// 		},
		// 	},
		// 	expected: `[{"tax":0},{"tax":0},{"tax":0},{"tax":10000}]`,
		// },
		// {
		// 	name: "Case 6",
		// 	input: []entity.Transaction{
		// 		{
		// 			Operation: "buy",
		// 			UnitCost:  10.00,
		// 			Quantity:  10000,
		// 		},
		// 		{
		// 			Operation: "sell",
		// 			UnitCost:  2.00,
		// 			Quantity:  5000,
		// 		},
		// 		{
		// 			Operation: "sell",
		// 			UnitCost:  20.00,
		// 			Quantity:  2000,
		// 		},
		// 		{
		// 			Operation: "sell",
		// 			UnitCost:  20.00,
		// 			Quantity:  2000,
		// 		},
		// 		{
		// 			Operation: "sell",
		// 			UnitCost:  25.00,
		// 			Quantity:  1000,
		// 		},
		// 	},
		// 	expected: `[{"tax":0},{"tax":0},{"tax":0},{"tax":3000}]`,
		// },

	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := ProcessTransactions(c.input)
			if got != c.expected {
				t.Errorf("expected %s, got %s", c.expected, got)
			}
		})
	}
}

func TestCalculateProfit(t *testing.T) {
	type values struct {
		currentWeightedAverage float64
		unitCost               float64
		quantity               float64
	}

	cases := []struct {
		name     string
		input    values
		expected float64
	}{
		{
			name:     "Profit",
			input:    values{currentWeightedAverage: 10.00, unitCost: 20.00, quantity: 100},
			expected: 1000.00,
		},
		{
			name:     "Loss",
			input:    values{currentWeightedAverage: 25.00, unitCost: 20.00, quantity: 5000},
			expected: -25000.00,
		},
		{
			name:     "Neutral",
			input:    values{currentWeightedAverage: 20.00, unitCost: 20.00, quantity: 10000},
			expected: 0.00,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := calculateProfit(c.input.currentWeightedAverage, c.input.unitCost, c.input.quantity)
			if got != c.expected {
				t.Errorf("expected %f, got %f", c.expected, got)
			}
		})
	}
}

func TestCalcWeightedAverage(t *testing.T) {
	type values struct {
		currentStocks          float64
		stocksBought           float64
		stocksCost             float64
		currentWeightedAverage float64
	}

	cases := []struct {
		name string
		input values
		expected float64
	}{
		{
			name: "Case 1",
			input: values{
				currentStocks:          5000,
				stocksBought:           3000,
				stocksCost:             8.00,
				currentWeightedAverage: 5.00,
			},
			expected: 6.13,
		},
		{
			name: "Case 2",
			input: values{
				currentStocks:          2000,
				stocksBought:           10000,
				stocksCost:             10.00,
				currentWeightedAverage: 15.00,
			},
			expected: 10.83,
		},
		{
			name: "Case 3",
			input: values{
				currentStocks:          0,
				stocksBought:           1000,
				stocksCost:             10.00,
				currentWeightedAverage: 0.00,
			},
			expected: 10.00,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := calcWeightedAverage(c.input.currentStocks, c.input.stocksBought, c.input.currentWeightedAverage, c.input.stocksCost)
			if got != c.expected {
				t.Errorf("expected %f, got %f", c.expected, got)
			}
		})
	}
}

func TestCalculateTax(t *testing.T) {
	type values struct {
		currentWeightedAverage float64
		profit                float64
		newProfit             float64
		in                   entity.Transaction
	}

	cases := []struct {
		name     string
		input    values
		expected float64
	}{
		{
			name: "Case 1 - shouldn't pay tax when Operation is buy",
			input: values{
				currentWeightedAverage: 10.00,
				profit:                0.00,
				newProfit:             1000.00,
				in: entity.Transaction{
					Operation: "buy",
					UnitCost:  20.00,
					Quantity:  1000,
				},
			},
			expected: 0.00,
		},
		{
			name: "Case 2 - shouldn't pay tax when operationCost is equal to 20000",
			input: values{
				currentWeightedAverage: 10.00,
				profit:                10.00,
				newProfit:             1000.00,
				in: entity.Transaction{
					Operation: "sell",
					UnitCost:  20.00,
					Quantity:  1000,
				},
			},
			expected: 0.00,
		},
		{
			name: "Case 3 - shouldn't pay tax when operationCost is equal to 20000",
			input: values{
				currentWeightedAverage: 10.00,
				profit:                10.00,
				newProfit:             1000.00,
				in: entity.Transaction{
					Operation: "sell",
					UnitCost:  20.00,
					Quantity:  1000,
				},
			},
			expected: 0.00,
		},
		{
			name: "Case 3 - shouldn't pay tax when operationCost is less than 20000",
			input: values{
				currentWeightedAverage: 10.00,
				profit:                30.00,
				newProfit:             20.00,
				in: entity.Transaction{
					Operation: "sell",
					UnitCost:  15.00,
					Quantity:  1000,
				},
			},
			expected: 0.00,
		},
		{
			name: "Case 4 - shouldn't pay tax when profit is equal to 0",
			input: values{
				currentWeightedAverage: 10.00,
				profit:                -20.00,
				newProfit:             20.00,
				in: entity.Transaction{
					Operation: "sell",
					UnitCost:  15.00,
					Quantity:  2000,
				},
			},
			expected: 0.00,
		},
		{
			name: "Case 5 - shouldn't pay tax when profit is less than 0",
			input: values{
				currentWeightedAverage: 10.00,
				profit:                -50.00,
				newProfit:             20.00,
				in: entity.Transaction{
					Operation: "sell",
					UnitCost:  15.00,
					Quantity:  2000,
				},
			},
			expected: 0.00,
		},
		{
			name: "Case 6 - should pay tax",
			input: values{
				currentWeightedAverage: 10.00,
				profit:                50.00,
				newProfit:             20.00,
				in: entity.Transaction{
					Operation: "sell",
					UnitCost:  15.00,
					Quantity:  2000,
				},
			},
			expected: 14.00,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, _ := calculateTax(c.input.currentWeightedAverage, c.input.profit, c.input.newProfit, c.input.in)
			if got != c.expected {
				t.Errorf("expected %f, got %f", c.expected, got)
			}
		})
	}
}