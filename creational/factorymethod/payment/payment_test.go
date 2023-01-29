package factorymethod

import (
	"testing"
)

func TestGetPayment(t *testing.T) {
	t.Run("Cash", func(t *testing.T) {
		want := CashPayment{Payment{"Huy", 10000000}}
		result := GetPayment(Cash, "Huy", 10000000)

		switch result := result.(type) {
		case *CashPayment:
			got := *result
			if got != want {
				t.Errorf("GetPayment(%d) = %v, want %v", Cash, got, want)
			}
		default:
			t.Errorf("GetPayment(%d) = %#v, want %#v", Cash, result, want)
		}
	})

	t.Run("Debit card", func(t *testing.T) {
		want := DebitCardPayment{Payment{"Huy", 10000000}}
		result := GetPayment(DebitCard, "Huy", 10000000)

		switch result := result.(type) {
		case *DebitCardPayment:
			got := *result
			if got != want {
				t.Errorf("GetPayment(%d) = %v, want %v", DebitCard, got, want)
			}
		default:
			t.Errorf("GetPayment(%d) = %#v, want %#v", DebitCard, result, want)
		}
	})
}
