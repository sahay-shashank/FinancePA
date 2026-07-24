package salary

import (
	"testing"
)

func TestNewSalaryInitializesEarningsForAllIncomeTypes(t *testing.T) {
	tests := []struct {
		name            string
		annualIncome    float64
		currency        string
		workHoursPerDay float64
		workDaysPerWeek int8
		incomeType      IncomeType
		wantType        IncomeType
		wantMonthly     float64
		wantWeekly      float64
		wantDaily       float64
		wantHourly      float64
	}{
		{
			name:            "salaried income uses default values",
			annualIncome:    120000,
			currency:        "INR",
			workHoursPerDay: 8,
			workDaysPerWeek: 5,
			incomeType:      "",
			wantType:        IncomeTypeSalaried,
			wantMonthly:     10000,
			wantWeekly:      2307.69,
			wantDaily:       461.54,
			wantHourly:      57.69,
		},
		{
			name:            "salaried income preserves the selected type",
			annualIncome:    120000,
			currency:        "INR",
			workHoursPerDay: 8,
			workDaysPerWeek: 5,
			incomeType:      IncomeTypeSalaried,
			wantType:        IncomeTypeSalaried,
			wantMonthly:     10000,
			wantWeekly:      2307.69,
			wantDaily:       461.54,
			wantHourly:      57.69,
		},
		{
			name:            "freelancer income preserves the selected type",
			annualIncome:    75000,
			currency:        "USD",
			workHoursPerDay: 6,
			workDaysPerWeek: 5,
			incomeType:      IncomeTypeFreelancer,
			wantType:        IncomeTypeFreelancer,
			wantMonthly:     6250,
			wantWeekly:      1442.31,
			wantDaily:       288.46,
			wantHourly:      48.08,
		},
		{
			name:            "business income preserves the selected type",
			annualIncome:    180000,
			currency:        "EUR",
			workHoursPerDay: 10,
			workDaysPerWeek: 6,
			incomeType:      IncomeTypeBusiness,
			wantType:        IncomeTypeBusiness,
			wantMonthly:     15000,
			wantWeekly:      3461.54,
			wantDaily:       576.92,
			wantHourly:      57.69,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSalary(tt.annualIncome, tt.currency, tt.workHoursPerDay, tt.workDaysPerWeek, tt.incomeType)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if got.IncomeType != tt.wantType {
				t.Fatalf("expected income type %q, got %q", tt.wantType, got.IncomeType)
			}

			assertFloat64(t, "monthly", got.Earnings.Monthly, tt.wantMonthly)
			assertFloat64(t, "weekly", got.Earnings.Weekly, tt.wantWeekly)
			assertFloat64(t, "daily", got.Earnings.Daily, tt.wantDaily)
			assertFloat64(t, "hourly", got.Earnings.Hourly, tt.wantHourly)
		})
	}
}

func TestNewSalaryReturnsEmptyEarningsForNonPositiveIncome(t *testing.T) {
	got, err := NewSalary(0, "INR", 8, 5, IncomeTypeFreelancer)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got.AnnualIncome != 0 {
		t.Fatalf("expected annual income 0, got %.2f", got.AnnualIncome)
	}

	if got.IncomeType != IncomeTypeFreelancer {
		t.Fatalf("expected income type %q, got %q", IncomeTypeFreelancer, got.IncomeType)
	}

	if got.Earnings.Monthly != 0 || got.Earnings.Weekly != 0 || got.Earnings.Daily != 0 || got.Earnings.Hourly != 0 {
		t.Fatalf("expected zero earnings for non-positive income, got %+v", got.Earnings)
	}
}

func TestNewSalaryReturnsErrorWhenCurrencyIsMissing(t *testing.T) {
	got, err := NewSalary(120000, "", 8, 5, IncomeTypeSalaried)
	if err == nil {
		t.Fatal("expected an error when currency is missing")
	}

	if got != nil {
		t.Fatalf("expected nil salary when currency is missing, got %+v", got)
	}
}

func TestSalaryMutatorsRecalculateEarnings(t *testing.T) {
	got, err := NewSalary(120000, "INR", 8, 5, IncomeTypeSalaried)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	got.SetAnnualIncome(240000)
	assertFloat64(t, "monthly", got.Earnings.Monthly, 20000)
	assertFloat64(t, "weekly", got.Earnings.Weekly, 4615.38)

	got.SetWorkDaysPerWeek(4)
	assertFloat64(t, "daily", got.Earnings.Daily, 1153.85)

	got.SetWorkHoursPerDay(10)
	assertFloat64(t, "hourly", got.Earnings.Hourly, 115.39)
}

func TestSalaryMutatorsUpdateIncomeTypeAndCurrency(t *testing.T) {
	got, err := NewSalary(120000, "INR", 8, 5, IncomeTypeSalaried)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	got.SetIncomeType(IncomeTypeFreelancer)
	if got.IncomeType != IncomeTypeFreelancer {
		t.Fatalf("expected income type %q, got %q", IncomeTypeFreelancer, got.IncomeType)
	}

	if err := got.SetCurrency("USD"); err != nil {
		t.Fatalf("expected no currency error, got %v", err)
	}

	if got.Currency != "USD" {
		t.Fatalf("expected currency %q, got %q", "USD", got.Currency)
	}
}

func TestSalarySetCurrencyRejectsEmptyCurrencyForZeroIncome(t *testing.T) {
	got := &Salary{}

	err := got.SetCurrency("")
	if err == nil {
		t.Fatal("expected an error when setting an empty currency for zero income")
	}
}

func TestSalaryMutatorResetsEarningsForZeroIncome(t *testing.T) {
	got, err := NewSalary(120000, "INR", 8, 5, IncomeTypeSalaried)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	got.SetAnnualIncome(0)
	if got.Earnings.Monthly != 0 || got.Earnings.Weekly != 0 || got.Earnings.Daily != 0 || got.Earnings.Hourly != 0 {
		t.Fatalf("expected zero earnings after resetting annual income, got %+v", got.Earnings)
	}
}

func assertFloat64(t *testing.T, label string, got float64, want float64) {
	t.Helper()
	if got != want {
		t.Fatalf("expected %s earnings %.2f, got %.2f", label, want, got)
	}
}
