package utils

import (
	"errors"
	"strings"
	"testing"

	"FinancePA/app/core"
	"FinancePA/app/utils"
)

// TestParseFromCSV ensures parsing goes as expected by perfomring parsing over different flavours of CSV.
func TestParseFromCSV(t *testing.T) {
	tests := []struct {
		name         string
		csvData      string
		expectedRows int
		expectedCols int
		expectError  bool
		expectedCode core.ErrorCode
	}{
		{
			name:         "Success: Valid standard financial statement records",
			csvData:      "2026-07-16,450.00,Groceries\n2026-07-17,1200.00,Homelab Server",
			expectedRows: 2,
			expectedCols: 3,
			expectError:  false,
		},
		{
			name:         "Success: Single cell spreadsheet data",
			csvData:      "SingleValue",
			expectedRows: 1,
			expectedCols: 1,
			expectError:  false,
		},
		{
			name:         "Failure: Malformed syntax quotes trigger parse error",
			csvData:      `2026-07-16,"Broken quote token parsing,Groceries`,
			expectError:  true,
			expectedCode: utils.CSVParseError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.csvData)

			records, err := utils.ParseFromCSV(reader)

			if tt.expectError {
				if err == nil {
					t.Fatalf("FAIL [%s]: Expected a parser engine crash, but received zero errors.", tt.name)
				}

				var appErr *core.AppError
				if errors.As(err, &appErr) {
					if appErr.Code != tt.expectedCode {
						t.Errorf("FAIL [%s]: Error signature mismatch.\nExpected Code: %s\nActual Code:   %s", tt.name, tt.expectedCode, appErr.Code)
					}
				} else {
					t.Errorf("FAIL [%s]: Error returned did not match our structured app/core.AppError type wrapper.", tt.name)
				}
				return
			}

			if err != nil {
				t.Fatalf("FAIL [%s]: Unexpected runtime parser error encountered: %v", tt.name, err)
			}

			if len(records) != tt.expectedRows {
				t.Errorf("FAIL [%s]: Total extracted data rows mismatch.\nExpected: %d\nActual:   %d", tt.name, tt.expectedRows, len(records))
			}

			if len(records) > 0 && len(records[0]) != tt.expectedCols {
				t.Errorf("FAIL [%s]: Row data column structural length mismatch.\nExpected: %d\nActual:   %d", tt.name, tt.expectedCols, len(records[0]))
			}
		})
	}
}
