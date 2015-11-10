package warehouse

import "fmt"

type Product int

//Some great products
const (
	BLAB Product = iota
	CAKE
	SUPERBLAB
)

type Order struct {
	quantity int
	product  Product
}

type Warehouse struct {
	inventory    int
	currentOrder Order
}

//Sets the warehouse's magical global inventory
func (w *Warehouse) SetInventory(qty int) error {
	w.inventory = qty
	return nil
}

func (w *Warehouse) PlaceOrder(product Product, qty int) error {
	if product.String() == "UNKNOWN" {
		return fmt.Errorf("Unknown Product specified")
	}

	if w.inventory >= qty {
		w.inventory -= qty
	}

	w.currentOrder = Order{qty, product}

	return nil
}

func (p Product) String() string {
	switch p {
	case BLAB:
		return fmt.Sprint("BLAB")
	case CAKE:
		return fmt.Sprint("CAKE")
	case SUPERBLAB:
		return fmt.Sprint("SUPERBLAB")
	default:
		return "UNKNOWN"
	}
}
