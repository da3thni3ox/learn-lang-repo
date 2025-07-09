package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	csvFile, err := os.Open("csv/test.csv")

	if err != nil {
		fmt.Println(err)
	}
	r := csv.NewReader(csvFile)

	records, err := r.ReadAll()

	if err != nil {
		fmt.Println(err)
		return
	}

	// Доработать как минимум миханизм добавления разделитиля
	for rowIdx, col := range records {
		for colIdx, row := range col {
			if rowIdx == 0 {
				fmt.Printf("%s", row)
				if colIdx == 2 {
					fmt.Println()
				}
			} else {
				fmt.Printf("%s", row)
				if colIdx == 2 {
					fmt.Println()
				}
			}
		}
	}
}
