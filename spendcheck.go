package spendcheck

import "errors"

var (
	// ErrInvalidExpenseAllocations is returned when the sum of the expense allocations is not 1.0.
	ErrInvalidExpenseAllocations = errors.New("invalid expense allocations")
)

// TransactionType is the type of transaction.
type TransactionType int

const (
	FixedCost TransactionType = iota
	Investment
	Savings
	Discretionary
)

// Transaction is a transaction.
type Transaction struct {
	Amount float64
	Type   TransactionType
}

// Plan is a budget plan.
type Plan struct {
	NetIncome     float64
	FixedCosts    float64
	Investments   float64
	Savings       float64
	Discretionary float64
	Transactions  []Transaction
}

// NewPlan creates a new Plan.
func NewPlan(netIncome, fixedCosts, investments, savings, discretionary float64) (*Plan, error) {
	totalAllocation := fixedCosts + investments + savings + discretionary
	if totalAllocation != 1.0 {
		return nil, ErrInvalidExpenseAllocations
	}
	return &Plan{
		NetIncome:     netIncome,
		FixedCosts:    fixedCosts,
		Investments:   investments,
		Savings:       savings,
		Discretionary: discretionary,
		Transactions:  []Transaction{},
	}, nil
}

// AddFixedCost adds a fixed cost transaction to the plan.
func (p *Plan) AddFixedCost(amount float64) {
	t := Transaction{
		Amount: amount,
		Type:   FixedCost,
	}
	p.Transactions = append(p.Transactions, t)
}

// AddInvestment adds an investment transaction to the plan.
func (p *Plan) AddInvestment(amount float64) {
	t := Transaction{
		Amount: amount,
		Type:   Investment,
	}
	p.Transactions = append(p.Transactions, t)
}

// AddSavings adds a savings transaction to the plan.
func (p *Plan) AddSavings(amount float64) {
	t := Transaction{
		Amount: amount,
		Type:   Savings,
	}
	p.Transactions = append(p.Transactions, t)
}

// AddDiscretionary adds a discretionary transaction to the plan.
func (p *Plan) AddDiscretionary(amount float64) {
	t := Transaction{
		Amount: amount,
		Type:   Discretionary,
	}
	p.Transactions = append(p.Transactions, t)
}

// SpendingSummary is a summary of spending.
type SpendingSummary struct {
	FixedCosts    float64
	Investments   float64
	Savings       float64
	Discretionary float64
}

// SummarizeSpending computes the spending among each category
// as a percentage of net income.
func (p *Plan) SummarizeSpending() SpendingSummary {
	var fixedCosts, investments, savings, discretionary float64
	for _, t := range p.Transactions {
		switch t.Type {
		case FixedCost:
			fixedCosts += t.Amount
		case Investment:
			investments += t.Amount
		case Savings:
			savings += t.Amount
		case Discretionary:
			discretionary += t.Amount
		}
	}
	return SpendingSummary{
		FixedCosts:    fixedCosts / p.NetIncome,
		Investments:   investments / p.NetIncome,
		Savings:       savings / p.NetIncome,
		Discretionary: discretionary / p.NetIncome,
	}
}
