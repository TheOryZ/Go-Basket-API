package helpers

import (
	"encoding/csv"
	"os"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/dtos"
)

//ReadToCsv reads a csv file and returns a slice of strings
func ReadToCsv(filePath string) (*[]dtos.CategoryCreateDTO, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.TrimLeadingSpace = true
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	categoryList := []dtos.CategoryCreateDTO{}
	for _, v := range records {
		categoryList = append(categoryList, dtos.CategoryCreateDTO{
			Name: v[0],
		})
	}
	return &categoryList, nil
}

//CreateCsvFile creates a csv file
func CreateCsvFile(filePath string, records [][]string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'

	err = writer.WriteAll(records)
	if err != nil {
		return err
	}

	writer.Flush()

	return nil
}
