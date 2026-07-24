package utils

import (
	"errors"
	"testing"
)

const (
	ErrInternalError    ErrorCode = "INTERNAL_SERVER_ERROR"
	ErrTelemetryFailure ErrorCode = "ERROR_TELEMETRY_FAILURE"
	ErrInvalidCSV       ErrorCode = "ERROR_INVALID_CSV"
)

// TestAppErrorFormatting uses Table-Driven Testing to validate error serialization
func TestAppErrorFormatting(t *testing.T) {
	tests := []struct {
		name           string
		inputError     *AppError
		expectedString string
	}{
		{
			name: "Simple error without underlying system error",
			inputError: &AppError{
				Code:    ErrInvalidCSV,
				Message: "Missing columns",
				Err:     nil,
			},
			expectedString: "[ERROR_INVALID_CSV] Missing columns",
		},
		{
			name: "Complex error wrapping an underlying OS/DB failure",
			inputError: &AppError{
				Code:    ErrTelemetryFailure,
				Message: "Network timed out for prometheus",
				Err:     errors.New("connection refused on port 9090"),
			},
			expectedString: "[ERROR_TELEMETRY_FAILURE] Network timed out for prometheus: connection refused on port 9090",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualString := tt.inputError.Error()

			if actualString != tt.expectedString {
				t.Errorf("\nFAIL: %s\nExpected: %q\nActual:   %q", tt.name, tt.expectedString, actualString)
			}
		})
	}
}

// TestAppErrorUnwrapping ensures our custom error can be unwrapped via standard Go tools
func TestAppErrorUnwrapping(t *testing.T) {
	underlyingOS := errors.New("disk full")
	customAppErr := NewAppError(ErrInternalError, "Failed to write dump block", underlyingOS)

	if !errors.Is(customAppErr, underlyingOS) {
		t.Errorf("FAIL: The underlying system error was lost during AppError wrapping chain.")
	}
}
