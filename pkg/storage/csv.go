package storage

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/juajosserand/goweb-challenge/internal/domain"
)

func ReadCSV(path string, dest *[]domain.Ticket) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}

	reader := csv.NewReader(f)
	data, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("could not read file: %w", err)
	}

	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return err
		}

		*dest = append(*dest, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}

	return nil
}
