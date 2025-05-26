package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/me/capital-gain/entity"
	"github.com/me/capital-gain/service"
)

func main() {
	var (
		inputBuilder strings.Builder
		// transactions []entity.Transaction
	)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputBuilder.WriteString(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		return
	}

	var transactions [][]entity.Transaction

	err := json.Unmarshal([]byte(inputBuilder.String()), &transactions)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error unmarshalling JSON:", err)
		return
	}

	output := service.ProcessTransactions(transactions)

	fmt.Println(output)
}
