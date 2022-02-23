package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {

	total := 0
	for _, v := range birdsPerDay {
		total += v
	}

	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {

	startIndex := (week * 7) - 7
	endIndex := startIndex + 7
	slice := birdsPerDay[startIndex:endIndex]

	return TotalBirdCount(slice)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {

	for k := range birdsPerDay {

		if k%2 == 0 {
			birdsPerDay[k] += 1
		}
	}
	return birdsPerDay
}
