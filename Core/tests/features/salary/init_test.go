package salary

import (
	"testing"

	"FinancePA/core/features/salary"
)

func TestNewSalaryCalculatesEarningsForAllIncomeTypes(t *testing.T) {
	tests := []struct {
		name            string
		annualIncome    float64
		workHoursPerDay float64
		workDaysPerWeek int8
		incomeType      salary.IncomeType
		expectedType    salary.IncomeType
		expectedMonthly float64
		expectedWeekly  float64
		expectedDaily   float64
		expectedHourly  float64
	}{
		{
			name:            "salaried income uses salaried defaults",
			annualIncome:    120000,
			workHoursPerDay: 8,
			workDaysPerWeek: 5,
			incomeType:      "",
			expectedType:    salary.IncomeTypeSalaried,
			expectedMonthly: 10000,
			expectedWeekly:  2307.69,
			expectedDaily:   461.54,
			expectedHourly:  57.69,
		}, {
			name:            "salaried income preserves the selected type",
			annualIncome:    120000,
			workHoursPerDay: 8,
			workDaysPerWeek: 5,
			incomeType:      salary.IncomeTypeSalaried,
			expectedType:    salary.IncomeTypeSalaried,
			expectedMonthly: 10000,
			expectedWeekly:  2307.69,
			expectedDaily:   461.54,
			expectedHourly:  57.69,
		},
		{
			name:            "freelancer income preserves the selected type",
			annualIncome:    75000,
			workHoursPerDay: 6,
			workDaysPerWeek: 5,
			incomeType:      salary.IncomeTypeFreelancer,
			expectedType:    salary.IncomeTypeFreelancer,
			expectedMonthly: 6250,
			expectedWeekly:  1442.31,
			expectedDaily:   288.46,
			expectedHourly:  48.08,
		},
		{
			name:            "business income preserves the selected type",
			annualIncome:    180000,
			workHoursPerDay: 10,
			workDaysPerWeek: 6,
			incomeType:      salary.IncomeTypeBusiness,
			expectedType:    salary.IncomeTypeBusiness,
			expectedMonthly: 15000,
			expectedWeekly:  3461.54,
			expectedDaily:   576.92,
			expectedHourly:  57.69,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := salary.NewSalary(tt.annualIncome, tt.workHoursPerDay, tt.workDaysPerWeek, tt.incomeType)

			if got.IncomeType != tt.expectedType {
				t.Fatalf("expected income type %q, got %q", tt.expectedType, got.IncomeType)
			}

			assertFloat64(t, "monthly", got.Earnings.Monthly, tt.expectedMonthly)
			assertFloat64(t, "weekly", got.Earnings.Weekly, tt.expectedWeekly)
			assertFloat64(t, "daily", got.Earnings.Daily, tt.expectedDaily)
			assertFloat64(t, "hourly", got.Earnings.Hourly, tt.expectedHourly)
		})
	}
}

func TestNewSalaryReturnsEmptyEarningsForNonPositiveIncome(t *testing.T) {
	got := salary.NewSalary(0, 8, 5, salary.IncomeTypeFreelancer)

	if got.AnnualIncome != 0 {
		t.Fatalf("expected annual income 0, got %.2f", got.AnnualIncome)
	}

	if got.IncomeType != salary.IncomeTypeFreelancer {
		t.Fatalf("expected income type %q, got %q", salary.IncomeTypeFreelancer, got.IncomeType)
	}

	if got.Earnings.Monthly != 0 || got.Earnings.Weekly != 0 || got.Earnings.Daily != 0 || got.Earnings.Hourly != 0 {
		t.Fatalf("expected zero earnings for non-positive income, got %+v", got.Earnings)
	}
}

func assertFloat64(t *testing.T, label string, got float64, want float64) {
	t.Helper()
	if got != want {
		t.Fatalf("expected %s earnings %.2f, got %.2f", label, want, got)
	}
}
