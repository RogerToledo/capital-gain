package main

import (
	"fmt"

	"github.com/me/capital-gain/entity"
	"github.com/me/capital-gain/service"
)

func main() {
	var transactions []entity.Transaction

	inputs, err := fmt.Scanf("%s", &transactions)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("Inputs:", inputs)

	output := service.ProcessTransactions(transactions)
	fmt.Println(output)
}
