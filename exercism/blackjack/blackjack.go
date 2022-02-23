package blackjack

var blackjackDict = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
	"other": 0,
}

const (
	stand = "S"
	hit   = "H"
	split = "P"
	win   = "W"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return blackjackDict[card]
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	return ParseCard(card1)+ParseCard(card2) == 21
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	switch {
	case isBlackjack &&
		dealerScore != 11 &&
		dealerScore != 10:
		return win
	case isBlackjack &&
		(dealerScore == 10 || dealerScore == 11):
		return stand
	case dealerScore < 20 && dealerScore != 10 && dealerScore >= 17:
		return win
	default:
		return split
	}
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {

	switch {
	case handScore <= 11:
		return hit
	case dealerScore >= 7 && (handScore >= 12 && handScore <= 16):
		return hit
	default:
		return stand
	}
}
