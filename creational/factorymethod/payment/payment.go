package factorymethod

type PaymentMethod interface {
	Pay(int)
}

type Payment struct {
	account string
	budget  int
}

func (p *Payment) Pay(amount int) {
}

type CashPayment struct {
	Payment
}

type DebitCardPayment struct {
	Payment
}

func newCashPayment(account string, budget int) PaymentMethod {
	return &CashPayment{
		Payment{
			account: account,
			budget:  budget,
		},
	}
}

func newDebitCardPayment(account string, budget int) PaymentMethod {
	return &DebitCardPayment{
		Payment{
			account: account,
			budget:  budget,
		},
	}
}

type PaymentType int

const (
	Cash PaymentType = iota
	DebitCard
)

func GetPayment(pmType PaymentType, account string, budget int) PaymentMethod {
	switch pmType {
	case Cash:
		return newCashPayment(account, budget)
	case DebitCard:
		return newDebitCardPayment(account, budget)
	default:
		return nil
	}
}
