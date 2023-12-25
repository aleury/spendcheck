package spendcheck_test

import (
	"spendcheck"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewPlan(t *testing.T) {
	t.Parallel()
	wantPlan := spendcheck.Plan{
		NetIncome:     1000.0,
		FixedCosts:    0.5,
		Investments:   0.1,
		Savings:       0.1,
		Discretionary: 0.3,
		Transactions:  []spendcheck.Transaction{},
	}
	p, err := spendcheck.NewPlan(1000.0, 0.5, 0.1, 0.1, 0.3)
	if err != nil {
		t.Fatal("didn't expect an error. got: ", err)
	}
	if !cmp.Equal(wantPlan, *p) {
		t.Error(cmp.Diff(wantPlan, *p))
	}
}

func TestNewPlan_ReturnsErrorForInvalidExpenseAllocations(t *testing.T) {
	t.Parallel()
	_, err := spendcheck.NewPlan(1000.0, 0.5, 0.1, 0.1, 0.4)
	if err == nil {
		t.Error("want error, got nil")
	}
}

func TestAddFixedCost(t *testing.T) {
	t.Parallel()
	p, err := spendcheck.NewPlan(1000.0, 0.5, 0.1, 0.1, 0.3)
	if err != nil {
		t.Fatal("didn't expect an error. got: ", err)
	}

	p.AddFixedCost(100.0)
	wantTransactions := []spendcheck.Transaction{
		{Amount: 100.0, Type: spendcheck.FixedCost},
	}
	if !cmp.Equal(wantTransactions, p.Transactions) {
		t.Error(cmp.Diff(wantTransactions, p.Transactions))
	}
}

func TestAddInvestment(t *testing.T) {
	t.Parallel()
	p, err := spendcheck.NewPlan(1000.0, 0.5, 0.1, 0.1, 0.3)
	if err != nil {
		t.Fatal("didn't expect an error. got: ", err)
	}

	p.AddInvestment(100.0)
	wantTransactions := []spendcheck.Transaction{
		{Amount: 100.0, Type: spendcheck.Investment},
	}
	if !cmp.Equal(wantTransactions, p.Transactions) {
		t.Error(cmp.Diff(wantTransactions, p.Transactions))
	}
}

func TestAddSavings(t *testing.T) {
	t.Parallel()
	p, err := spendcheck.NewPlan(1000.0, 0.5, 0.1, 0.1, 0.3)
	if err != nil {
		t.Fatal("didn't expect an error. got: ", err)
	}

	p.AddSavings(100.0)
	wantTransactions := []spendcheck.Transaction{
		{Amount: 100.0, Type: spendcheck.Savings},
	}
	if !cmp.Equal(wantTransactions, p.Transactions) {
		t.Error(cmp.Diff(wantTransactions, p.Transactions))
	}
}

func TestAddDiscretionary(t *testing.T) {
	t.Parallel()
	p, err := spendcheck.NewPlan(1000.0, 0.5, 0.1, 0.1, 0.3)
	if err != nil {
		t.Fatal("didn't expect an error. got: ", err)
	}

	p.AddDiscretionary(300.0)
	wantTransactions := []spendcheck.Transaction{
		{Amount: 300.0, Type: spendcheck.Discretionary},
	}
	if !cmp.Equal(wantTransactions, p.Transactions) {
		t.Error(cmp.Diff(wantTransactions, p.Transactions))
	}
}

func TestSummarizeSpending(t *testing.T) {
	t.Parallel()
	p, err := spendcheck.NewPlan(1000.0, 0.5, 0.1, 0.1, 0.3)
	if err != nil {
		t.Fatal("didn't expect an error. got: ", err)
	}

	p.AddFixedCost(500.0)
	p.AddInvestment(100.0)
	p.AddSavings(100.0)
	p.AddDiscretionary(250.0)
	summary := p.SummarizeSpending()
	wantSummary := spendcheck.SpendingSummary{
		FixedCosts:    .5,
		Investments:   .1,
		Savings:       .1,
		Discretionary: .25,
	}
	if !cmp.Equal(wantSummary, summary) {
		t.Error(cmp.Diff(wantSummary, summary))
	}
}
