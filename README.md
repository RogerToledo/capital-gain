# Capital Gain Tax Calculator
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/RogerToledo/capital-gain
)

## Objective

This project is part of a selection process step for Nubank.

## Description

This Go service processes a series of stock transactions (buy and sell operations) and calculates the applicable capital gains tax for each sell operation. It utilizes a weighted average cost method to determine profit and applies a 20% tax rate on profits exceeding a certain transaction volume.

## How to run the tests

`go test ./...`

## Inputs and outputs
The application receive a json as input with 3 keys:
- operation
- unit-cost
- quantity

Follow an example
```
[
    [
        {"operation":"buy", "unit-cost":10.00, "quantity": 10000},
        {"operation":"sell", "unit-cost":20.00, "quantity": 5000},
        {"operation":"sell", "unit-cost":5.00, "quantity": 5000}
    ]
]
```
And return a json with a key tax
```
[
    [
        {"tax": 0.0},{"tax": 10000.0},{"tax": 0.0}
    ]
]
```

## How to run the application
`go run main.go`

or

`cat input.json | go run main.go`

if you have a json file (input.json is an example)


## Technologies
- [Go](https://golang.org/)

