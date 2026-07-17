package utils

import (
	"FinancePA/core/utils"
	"testing"
)

func TestRoundToTwoDecimal(t *testing.T) {
	got := utils.RoundToTwoDecimal(2307.6923076923076)
	if got != 2307.69 {
		t.Fatalf("expected 2307.69, got %.2f", got)
	}
}
