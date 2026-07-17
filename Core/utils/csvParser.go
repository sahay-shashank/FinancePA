package utils

import (
	"encoding/csv"
	"io"
)

const (
	CSVParseError ErrorCode = "CSV_PARSE_ERROR"
)

func ParseFromCSV(readerIO io.Reader) ([][]string, error) {
	csvReader := csv.NewReader(readerIO)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, NewAppError(CSVParseError, "Unable to parse CSV", err)
	}
	return records, nil
}
