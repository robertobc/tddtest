package warehouse

import (
	"fmt"
	"testing"
)

func TestSetInventory(t *testing.T) {
	w := &Warehouse{}
	cases := []struct {
		qty         int
		expectedQty int
		expectedErr error
	}{
		{123, 123, nil},
		{88, 88, nil},
		{-23, 88, fmt.Errorf("Invalid qty supplied")},
	}

	for _, c := range cases {
		err := w.SetInventory(c.qty)

		if !compareErr(err, c.expectedErr) {
			t.Errorf("Expected a %v error, got a %v error", c.expectedErr, err)
		}

		if w.inventory != c.expectedQty {
			t.Errorf("Inventory expected: %d, Actual: %d", c.expectedQty, w.inventory)
		}
	}
}

func TestPlaceOrder(t *testing.T) {
	w := &Warehouse{}
	w.inventory = 100

	cases := []struct {
		product     Product
		qty         int
		expectedQty int
		expectedErr error
	}{
		{BLAB, 20, 80, nil},
		{SUPERBLAB, 30, 50, nil},
		{CAKE, 51, 50, fmt.Errorf("Not enough inventory to satisfy your order")},
		{123, 10, 50, fmt.Errorf("Unknown Product specified")},
		{BLAB, 0, 50, fmt.Errorf("Qty specified is invalid")},
		{CAKE, -55, 50, fmt.Errorf("Qty specified is invalid")},
	}

	for _, c := range cases {
		err := w.PlaceOrder(c.product, c.qty)

		if !compareErr(err, c.expectedErr) {
			t.Errorf("Expected a %v error, got a %v error", c.expectedErr, err)
		}

		if w.inventory != c.expectedQty {
			t.Errorf("Inventory expected: %d, Actual: %d", c.expectedQty, w.inventory)
		}
	}
}

func compareErr(err1, err2 error) bool {
	if err1 == err2 {
		return true
	}

	if err1 != nil && err2 != nil && err1.Error() == err2.Error() {
		return true
	}

	return false
}
