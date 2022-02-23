package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	switch kind {
	case "car":
		return true
	case "truck":
		return true
	default:
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {

	choice := ""
	if option1 < option2 {
		choice = option1
	} else {
		choice = option2
	}
	return choice + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {

	const (
		RvHigh   = .8
		RvLow    = .5
		RvMedium = .7
	)

	switch {

	case age < 3:
		return originalPrice * RvHigh
	case age > 3 && age < 10:
		return originalPrice * RvMedium
	default:
		return originalPrice * RvLow
	}

}
