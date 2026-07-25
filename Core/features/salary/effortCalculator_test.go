package salary

import (
	"testing"
)

func TestCalculateEffort(t *testing.T) {
	salary, err := NewSalary(120000, "INR", 8, 5, IncomeTypeSalaried)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	testCases := []struct {
		testCase       string
		amount         float64
		expectedYears  int
		expectedMonths int
		expectedWeeks  int
		expectedDays   int
		expectedHours  float32
	}{
		{
			testCase:       "no amount",
			amount:         0,
			expectedYears:  0,
			expectedMonths: 0,
			expectedWeeks:  0,
			expectedDays:   0,
			expectedHours:  0.0,
		},
		{
			testCase:       "negative amount",
			amount:         -5000,
			expectedYears:  0,
			expectedMonths: 0,
			expectedWeeks:  0,
			expectedDays:   0,
			expectedHours:  0.0,
		},
		{
			testCase:       "hours only",
			amount:         57.69,
			expectedYears:  0,
			expectedMonths: 0,
			expectedWeeks:  0,
			expectedDays:   0,
			expectedHours:  1.0,
		},
		{
			testCase:       "days only",
			amount:         923.08,
			expectedYears:  0,
			expectedMonths: 0,
			expectedWeeks:  0,
			expectedDays:   2,
			expectedHours:  0.0,
		},
		{
			testCase:       "weeks only",
			amount:         4615.38,
			expectedYears:  0,
			expectedMonths: 0,
			expectedWeeks:  2,
			expectedDays:   0,
			expectedHours:  0.0,
		},
		{
			testCase:       "months only",
			amount:         30000,
			expectedYears:  0,
			expectedMonths: 3,
			expectedWeeks:  0,
			expectedDays:   0,
			expectedHours:  0.0,
		},
		{
			testCase:       "years only",
			amount:         240000,
			expectedYears:  2,
			expectedMonths: 0,
			expectedWeeks:  0,
			expectedDays:   0,
			expectedHours:  0.0,
		},
		{
			testCase:       "complex time",
			amount:         132827.92,
			expectedYears:  1,
			expectedMonths: 1,
			expectedWeeks:  1,
			expectedDays:   1,
			expectedHours:  1.0,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testCase, func(t *testing.T) {
			got := salary.CalculateEffort(tt.amount)
			if got.Years != tt.expectedYears {
				t.Fatalf("expected years %d, got %d", tt.expectedYears, got.Years)
			}
			if got.Months != tt.expectedMonths {
				t.Fatalf("expected months %d, got %d", tt.expectedMonths, got.Months)
			}
			if got.Weeks != tt.expectedWeeks {
				t.Fatalf("expected weeks %d, got %d", tt.expectedWeeks, got.Weeks)
			}
			if got.Days != tt.expectedDays {
				t.Fatalf("expected days %d, got %d", tt.expectedDays, got.Days)
			}
			if !((tt.expectedHours-0.1) < got.Hours && got.Hours < (tt.expectedHours+0.1)) {
				t.Fatalf("expected hours ~%.2f (+-0.1), got %.2f", tt.expectedHours, got.Hours)
			}
		})
	}
}

func TestCalculateEffortWithZeroSalaryAndHighSpending(t *testing.T) {
	zeroSalary, err := NewSalary(0, "INR", 8, 5, IncomeTypeSalaried)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// When salary is 0, spending money is impossible
	// Should return zero Time (not infinity or garbage)
	got := zeroSalary.CalculateEffort(50000)

	if got.Years != 0 {
		t.Fatalf("expected years 0 for zero salary, got %d", got.Years)
	}
	if got.Months != 0 {
		t.Fatalf("expected months 0 for zero salary, got %d", got.Months)
	}
	if got.Weeks != 0 {
		t.Fatalf("expected weeks 0 for zero salary, got %d", got.Weeks)
	}
	if got.Days != 0 {
		t.Fatalf("expected days 0 for zero salary, got %d", got.Days)
	}
	if got.Hours != 0 {
		t.Fatalf("expected hours 0 for zero salary, got %f", got.Hours)
	}
}
