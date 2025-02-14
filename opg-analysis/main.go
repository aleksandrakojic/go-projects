package main

import (
	"fmt"
	"os"
)

type Stock struct {
	Ticker			string
	Gap				float64
	OpeningPrice	float64
}

func Load(path string) ([]Stock, error) {
	f, err := os.Open(path)
	if err != nil {
        fmt.Printf("Error opening", err)
		return nil, err
    }
	defer f.Close()

	r := csv.NewReader(f)
	rows, err := r.ReadAll()

	if err!= nil {
        fmt.Printf("Error reading CSV file", err)
        return nil, err
    }

	rows = slices.Delete(rows, 0, 1)
	var stocks []Stock

	for _, row := range rows {
        ticker := row()

		gap, err := strconv.ParseFloat(row[1], 64)
		if err!= nil {
            fmt.Printf("Error parsing gap: %v", err)
            continue
        }
		openingPrice, err := strconv.ParseFloat(row(2), 64)

		stocks = append(stocks, Stock{
			Ticker: ticker,
            Gap: atof(gap),
            OpeningPrice: atof(openingPrice),
    }

	return stocks, nil
}

func main() {
	stocks, err := Load("./opg.csv")
	if err!= nil {
        fmt.Println(err)
        return
    }

	slices.DeleteFunc(stocks, func(s Stock) bool) {
		return math.Abs(s.Gap) < .1
	}
}