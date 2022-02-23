package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minute int) int {
	if minute == 0 {
		return len(layers) * 2
	}

	return len(layers) * minute
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {

	noodles := 0
	sauce := 0.0

	for _, v := range layers {
		switch v {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		default:
		}
	}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	secret := friendsList[len(friendsList)-1:]
	myList[len(myList)-1] = secret[0]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, amount int) []float64 {

	numberToScale := float64(amount) / 2
	newQuantities := []float64{}

	for _, v := range quantities {
		newQuantities = append(newQuantities, v*numberToScale)
	}

	return newQuantities
}
