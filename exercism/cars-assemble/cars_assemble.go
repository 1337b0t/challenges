package cars

const GroupProductionRate = 95000
const IndividualProductionRate = 10000
const Group = 10

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return successRate / 100 * float64(productionRate)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {

	sr := successRate / 100
	pm := float64(productionRate) / 60
	result := int(sr * pm)

	return result
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {

	carGroup := carsCount / Group
	carRemainder := carsCount % Group

	return uint((GroupProductionRate * carGroup) + (IndividualProductionRate * carRemainder))
}
