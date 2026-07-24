package salary

import (
	"testing"
)

func TestNewSalaryNormalizesInvalidWorkSchedule(t *testing.T) {
	got, err := NewSalary(120000, "INR", 0, 0, IncomeTypeSalaried)

	if err != nil {
		t.Fatal("error creating salary object")
	}

	if got.WorkHoursPerDay != 8 {
		t.Fatalf("expected default work hours per day 8, got %.2f", got.WorkHoursPerDay)
	}

	if got.WorkDaysPerWeek != 5 {
		t.Fatalf("expected default work days per week 5, got %d", got.WorkDaysPerWeek)
	}
}
