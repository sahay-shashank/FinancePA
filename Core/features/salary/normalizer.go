package salary

func normalizeIncomeType(incomeType IncomeType) IncomeType {
	switch incomeType {
	case IncomeTypeFreelancer:
		return IncomeTypeFreelancer
	case IncomeTypeBusiness:
		return IncomeTypeBusiness
	default:
		return IncomeTypeSalaried
	}
}

func normalizeWorkDaysPerWeek(workDaysPerWeek int8) int8 {
	normalized := workDaysPerWeek
	if normalized <= 0 {
		normalized = 5
	}
	return normalized
}

func normalizeWorkHoursPerDay(workHoursPerDay float64) float64 {
	normalized := workHoursPerDay
	if normalized <= 0 {
		normalized = 8
	}
	return normalized
}
