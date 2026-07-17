package salary

import (
	"FinancePA/core/features/salary"
	"testing"
)

func TestNewSalaryNormalizesInvalidWorkSchedule(t *testing.T) {
	got := salary.NewSalary(120000, 0, 0, salary.IncomeTypeSalaried)

	if got.WorkHoursPerDay != 8 {
		t.Fatalf("expected default work hours per day 8, got %.2f", got.WorkHoursPerDay)
	}

	if got.WorkDaysPerWeek != 5 {
		t.Fatalf("expected default work days per week 5, got %d", got.WorkDaysPerWeek)
	}
}
