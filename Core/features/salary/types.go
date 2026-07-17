package salary

type IncomeType string

const (
	IncomeTypeSalaried   IncomeType = "salaried"
	IncomeTypeFreelancer IncomeType = "freelancer"
	IncomeTypeBusiness   IncomeType = "business"
)

type Salary struct {
	AnnualIncome    float64      `json:"annual_income"`
	WorkHoursPerDay float64      `json:"work_hours_per_day"`
	WorkDaysPerWeek int8         `json:"work_days_per_week"`
	IncomeType      IncomeType   `json:"income_type"`
	Earnings        timeEarnings `json:"earnings"`
}

type timeEarnings struct {
	Hourly  float64 `json:"hourly"`
	Daily   float64 `json:"daily"`
	Weekly  float64 `json:"weekly"`
	Monthly float64 `json:"monthly"`
}

const (
	monthsPerYear = 12
	weeksPerYear  = 52
)
