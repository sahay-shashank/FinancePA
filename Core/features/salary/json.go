package salary

import (
	"FinancePA/core/utils"
	"encoding/json"
)

const (
	SalaryJsonUnmarshalError utils.ErrorCode = "SALARY_JSON_UNMARSHAL_ERROR"
	SalaryJsonMarshalError   utils.ErrorCode = "SALARY_JSON_MARSHAL_ERROR"
)

var (
	jsonMarshal   = json.Marshal
	jsonUnMarshal = json.Unmarshal
)

type salaryPayload struct {
	AnnualIncome    float64      `json:"annual_income"`
	Currency        string       `json:"currency"`
	WorkHoursPerDay float64      `json:"work_hours_per_day"`
	WorkDaysPerWeek int8         `json:"work_days_per_week"`
	IncomeType      IncomeType   `json:"income_type"`
	Earnings        TimeEarnings `json:"earnings"`
}

func UnmarshalJSON(payload []byte) (*Salary, error) {
	salary := &Salary{}
	err := jsonUnMarshal(payload, salary)
	if err != nil {
		return nil, utils.NewAppError(SalaryJsonUnmarshalError, "Unable to unmarshal salary payload", err)
	}
	return salary, nil
}

func (salary *Salary) MarshalJSON() ([]byte, error) {
	payload, err := jsonMarshal(salaryPayload{
		AnnualIncome:    salary.AnnualIncome,
		Currency:        salary.Currency,
		WorkHoursPerDay: salary.WorkHoursPerDay,
		WorkDaysPerWeek: salary.WorkDaysPerWeek,
		IncomeType:      salary.IncomeType,
		Earnings:        salary.Earnings,
	})
	if err != nil {
		return nil, utils.NewAppError(SalaryJsonMarshalError, "Unable to marshal salary payload", err)
	}
	return payload, nil
}
