package salary

import "FinancePA/core/utils"

type IncomeType string

const (
	IncomeTypeSalaried   IncomeType = "salaried"
	IncomeTypeFreelancer IncomeType = "freelancer"
	IncomeTypeBusiness   IncomeType = "business"
	monthsPerYear                   = 12
	weeksPerYear                    = 52
)

// ERROR Constants
const (
	NoCurrency utils.ErrorCode = "NO_CURRENCY"
)

type Salary struct {
	AnnualIncome    float64      `json:"annual_income"`
	Currency        string       `json:"currency"`
	WorkHoursPerDay float64      `json:"work_hours_per_day"`
	WorkDaysPerWeek int8         `json:"work_days_per_week"`
	IncomeType      IncomeType   `json:"income_type"`
	Earnings        TimeEarnings `json:"earnings"`
}

type TimeEarnings struct {
	Hourly  float64 `json:"hourly"`
	Daily   float64 `json:"daily"`
	Weekly  float64 `json:"weekly"`
	Monthly float64 `json:"monthly"`
}

func NewSalary(annualIncome float64, currency string, workHoursPerDay float64, workDaysPerWeek int8, incomeType IncomeType) (*Salary, *utils.AppError) {
	incomeType = normalizeIncomeType(incomeType)

	if annualIncome <= 0 {
		return &Salary{IncomeType: incomeType}, nil
	}

	if currency == "" {
		return nil, utils.NewAppError(NoCurrency, "no currency declared for annual income", nil)
	}

	normalizedWorkDaysPerWeek := normalizeWorkDaysPerWeek(workDaysPerWeek)

	normalizedWorkHoursPerDay := normalizeWorkHoursPerDay(workHoursPerDay)

	salaryValue := &Salary{
		AnnualIncome:    annualIncome,
		Currency:        currency,
		WorkHoursPerDay: normalizedWorkHoursPerDay,
		WorkDaysPerWeek: normalizedWorkDaysPerWeek,
		IncomeType:      incomeType,
	}

	salaryValue.refreshEarnings()

	return salaryValue, nil
}

func (s *Salary) SetAnnualIncome(annualIncome float64) {
	s.AnnualIncome = annualIncome
	s.refreshEarnings()
}

func (s *Salary) SetWorkHoursPerDay(workHoursPerDay float64) {
	s.WorkHoursPerDay = normalizeWorkHoursPerDay(workHoursPerDay)
	s.refreshEarnings()
}

func (s *Salary) SetWorkDaysPerWeek(workDaysPerWeek int8) {
	s.WorkDaysPerWeek = normalizeWorkDaysPerWeek(workDaysPerWeek)
	s.refreshEarnings()
}

func (s *Salary) SetIncomeType(incomeType IncomeType) {
	s.IncomeType = normalizeIncomeType(incomeType)
	s.refreshEarnings()
}
func (s *Salary) SetCurrency(currency string) *utils.AppError {
	if s.AnnualIncome <= 0 && currency == "" {
		return utils.NewAppError(NoCurrency, "No currency declared for annual income", nil)
	}
	s.Currency = currency
	return nil
}

func (s *Salary) refreshEarnings() {
	if s.AnnualIncome <= 0 {
		s.Earnings = TimeEarnings{}
		return
	}

	monthlySalary := utils.RoundToTwoDecimal(s.AnnualIncome / monthsPerYear)
	weeklySalary := utils.RoundToTwoDecimal(s.AnnualIncome / weeksPerYear)
	dailySalary := utils.RoundToTwoDecimal(weeklySalary / float64(s.WorkDaysPerWeek))
	hourlySalary := utils.RoundToTwoDecimal(dailySalary / s.WorkHoursPerDay)

	s.Earnings = TimeEarnings{
		Monthly: monthlySalary,
		Weekly:  weeklySalary,
		Daily:   dailySalary,
		Hourly:  hourlySalary,
	}
}
