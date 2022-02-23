package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {

	val, ok := inMap(units, unit)
	if ok {
		bill[item] += val
	}

	return ok
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {

	billVal, ok := inMap(bill, item)

	if ok {
		itemVal, ok := inMap(units, unit)
		if ok {

			switch newQty := billVal - itemVal; {
			case newQty == 0:
				delete(bill, item)
			case newQty > 0:
				bill[item] = newQty
			default:
				return false
			}

			return true
		}
	}

	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	return inMap(bill, item)
}

func inMap(m map[string]int, item string) (int, bool) {

	if val, ok := m[item]; ok {
		return val, ok
	}
	return 0, false
}
