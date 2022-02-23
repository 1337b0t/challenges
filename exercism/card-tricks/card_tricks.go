package cards

func isValidIndex(slice []int, index int) bool {
	uBound := len(slice) - 1

	if index < 0 || index > uBound {
		return false
	}

	return true
}

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {

	if !isValidIndex(slice, index) {
		return 0, false
	}
	return slice[index], true
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {

	if !isValidIndex(slice, index) {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}

	return slice
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {

	slice := []int{}
	i := 1
	if length > 0 {
		for i <= length {
			slice = append(slice, value)
			i++
		}
	}

	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {

	if isValidIndex(slice, index) {
		slice = append(slice[:index], slice[index+1:]...)
	}

	return slice
}
