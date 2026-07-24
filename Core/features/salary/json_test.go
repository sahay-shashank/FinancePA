package salary

import (
	"encoding/json"
	"errors"
	"testing"

	"FinancePA/core/utils"
)

func TestMarshalJSON(t *testing.T) {
	s := &Salary{
		AnnualIncome:    120000,
		Currency:        "INR",
		WorkHoursPerDay: 8,
		WorkDaysPerWeek: 5,
		IncomeType:      IncomeTypeSalaried,
		Earnings: TimeEarnings{
			Monthly: 10000,
			Weekly:  2307.69,
			Daily:   461.54,
			Hourly:  57.69,
		},
	}

	payload, err := s.MarshalJSON()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var decoded map[string]any
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("expected payload to be valid JSON, got %v", err)
	}

	if decoded["annual_income"] != float64(120000) {
		t.Fatalf("expected annual_income to be serialized, got %#v", decoded["annual_income"])
	}
}

func TestUnmarshalJSONReturnsAppErrorOnFailure(t *testing.T) {
	payload := []byte(`{"annual_income":`)

	got, err := UnmarshalJSON(payload)
	if err == nil {
		t.Fatal("expected an error for malformed JSON")
	}
	if got != nil {
		t.Fatalf("expected nil salary on error, got %+v", got)
	}
	var appErr *utils.AppError
	if !errors.As(err, &appErr) {
		t.Fatalf("expected AppError, got %T", err)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	payload := []byte(`{"annual_income":120000,"currency":"INR","work_hours_per_day":8,"work_days_per_week":5,"income_type":"salaried","earnings":{"monthly":10000,"weekly":2307.69,"daily":461.54,"hourly":57.69}}`)

	got, err := UnmarshalJSON(payload)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got.AnnualIncome != 120000 {
		t.Fatalf("expected annual income 120000, got %.2f", got.AnnualIncome)
	}
	if got.Currency != "INR" {
		t.Fatalf("expected currency INR, got %q", got.Currency)
	}
	if got.IncomeType != IncomeTypeSalaried {
		t.Fatalf("expected income type %q, got %q", IncomeTypeSalaried, got.IncomeType)
	}
}

func TestMarshalJSONReturnsAppErrorOnFailure(t *testing.T) {
	salaryTest := &Salary{}
	original := jsonMarshal
	jsonMarshal = func(v any) ([]byte, error) {
		return nil, errors.New("marshal failed")
	}
	defer func() { jsonMarshal = original }()

	payload, err := salaryTest.MarshalJSON()
	if err == nil {
		t.Fatal("expected an error when marshaling fails")
	}
	if payload != nil {
		t.Fatalf("expected nil payload on error, got %v", payload)
	}

	var appErr *utils.AppError
	if !errors.As(err, &appErr) {
		t.Fatalf("expected AppError, got %T", err)
	}
}
