package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("rates.json")
	if err != nil {
		panic(err)
	}

	var data struct {
		Base  string             `json:"base"`
		Date  string             `json:"date"`
		Rates map[string]float64 `json:"rates"`
	}

	if err := json.Unmarshal(file, &data); err != nil {
		panic(err)
	}

	if len(os.Args) != 3 {
		fmt.Println("Erro: quantidade de argumentos inválida.")
		os.Exit(1)
	}

	value := os.Args[1]
	currency := strings.ToUpper(os.Args[2])

	valueBRL, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("Valor inválido para BRL.")
		os.Exit(1)
	}

	rate, ok := data.Rates[currency]
	if !ok {
		fmt.Printf("Erro: moeda %s não encontrada. \n", currency)
		os.Exit(1)
	}

	convertedValue := valueBRL * rate

	fmt.Printf("%.2f BRL = %.2f %s na cotação de %s e taxa = %.3f", valueBRL, convertedValue, currency, data.Date, rate)
}
