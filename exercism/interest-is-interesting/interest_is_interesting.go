package interest

func getInterest(balance float64) float64 {
	switch {
	case balance >= 0 && balance < 1000:
		return .5
	case balance >= 1000 && balance < 5000:
		return 1.621
	case balance >= 5000:
		return 2.475
	default:
		return 3.213
	}
}

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	return float32(getInterest(balance))
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return getInterest(balance) / 100.0 * balance
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var numOfYears int = 1
	for {
		balance = AnnualBalanceUpdate(balance)
		if balance >= targetBalance {
			break
		}
		numOfYears++
	}
	return numOfYears
}
