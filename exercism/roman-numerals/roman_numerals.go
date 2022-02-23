package romannumerals

import "errors"

type numeral struct {
	Num   int
	Value string
}

var (
	numerals = []numeral{
		0:  numeral{1000, "M"},
		1:  numeral{900, "CM"},
		2:  numeral{500, "D"},
		3:  numeral{400, "CD"},
		4:  numeral{100, "C"},
		5:  numeral{90, "XC"},
		6:  numeral{50, "L"},
		7:  numeral{40, "XL"},
		8:  numeral{10, "X"},
		9:  numeral{9, "IX"},
		10: numeral{5, "V"},
		11: numeral{4, "IV"},
		12: numeral{1, "I"},
	}
)

func ToRomanNumeral(input int) (string, error) {

	result := ""

	if input > 0 && input <= 3000 {

		for _, v := range numerals {
			for input >= v.Num {
				result += v.Value
				input -= v.Num
			}
		}

		return result, nil
	}

	return result, errors.New("number must be greater than 0")
}
