package salary

import "FinancePA/core/utils"

func NewSalary(annualIncome float64, workHoursPerDay float64, workDaysPerWeek int8, incomeType IncomeType) *Salary {
	incomeType = normalizeIncomeType(incomeType)

	if annualIncome <= 0 {
		return &Salary{IncomeType: incomeType}
	}

	normalizedWorkDaysPerWeek := normalizeWorkDaysPerWeek(workDaysPerWeek)

	normalizedWorkHoursPerDay := normalizeWorkHoursPerDay(workHoursPerDay)

	monthlySalary := utils.RoundToTwoDecimal(annualIncome / monthsPerYear)
	weeklySalary := utils.RoundToTwoDecimal(annualIncome / weeksPerYear)
	dailySalary := utils.RoundToTwoDecimal(weeklySalary / float64(normalizedWorkDaysPerWeek))
	hourlySalary := utils.RoundToTwoDecimal(dailySalary / normalizedWorkHoursPerDay)

	return &Salary{
		AnnualIncome:    annualIncome,
		WorkHoursPerDay: normalizedWorkHoursPerDay,
		WorkDaysPerWeek: normalizedWorkDaysPerWeek,
		IncomeType:      incomeType,
		Earnings: timeEarnings{
			Monthly: monthlySalary,
			Weekly:  weeklySalary,
			Daily:   dailySalary,
			Hourly:  hourlySalary,
		},
	}
}
