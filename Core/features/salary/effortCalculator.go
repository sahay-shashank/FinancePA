package salary

func (salary *Salary) CalculateEffort(amount float64) Time {
	var years int = 0
	var months int = 0
	var weeks int = 0
	var days int = 0
	var hours float32 = 0.0

	if amount > 0 && salary.AnnualIncome > 0 {
		// Calculate years
		years = int(amount / salary.AnnualIncome)
		amount = amount - float64(years)*salary.AnnualIncome

		// Calculate months
		monthlyIncome := salary.Earnings.Monthly
		months = int(amount / monthlyIncome)
		amount = amount - float64(months)*monthlyIncome

		// Calculate weeks
		weeklyIncome := salary.Earnings.Weekly
		weeks = int(amount / weeklyIncome)
		amount = amount - float64(weeks)*weeklyIncome

		// Calculate days
		dailyIncome := salary.Earnings.Daily
		days = int(amount / dailyIncome)
		amount = amount - float64(days)*dailyIncome

		// Calculate hours
		hourlyIncome := salary.Earnings.Hourly
		hours = float32(amount / hourlyIncome)
	}

	return Time{
		Years:  years,
		Months: months,
		Weeks:  weeks,
		Days:   days,
		Hours:  hours,
	}
}
