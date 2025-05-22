package main

import (
	"fmt"
	"math"
)

// [ ] TODO: Tax 20% sobre o lucro obtido na operação. O imposto só vai ser pago se quando há uma operação
// de venda cujo o preço é superior ao preço médio ponderado de compra

// [ ] TODO: Preço médio ponderado -> novaMediaPonderada = ((quantidadeAcoesAtual * mediaPonderadaAtual) +
// + (quantidadeAcoesCompradas * valorCompra)) / (quantidadeAcoesAtual + quantidadeAcoesCompradas)

func main() {
	currentStocks := 10000
	currentWeightedAverage := 10.0

	stocksCost := 20.0
	stocksBought := 10

	newWeightedAverage := calcWeightedAverage(currentStocks, stocksBought, currentWeightedAverage, stocksCost)

	fmt.Println("New Weighted Average:", newWeightedAverage)
}

func calcWeightedAverage(currentStocks, stocksBought int, currentWeightedAverage, stocksCost float64) float64 {
	wa := (float64(currentStocks) * currentWeightedAverage) + (float64(stocksBought) * stocksCost) / (float64(currentStocks) + float64(stocksBought))
	roundedWa := math.Round(wa*100) / 100
	
	return roundedWa
}
