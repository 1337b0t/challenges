package phonenumber

import (
	"errors"
	"fmt"
	"strings"
)

func Number(phoneNumber string) (string, error) {

	errorAreaCode := errors.New("1 or 0 are not allowed in the area code")
	errorExchangeCode := errors.New("1 and 0 are not allowed in the exchange code")
	errorLength := errors.New("the phone number length is not 10")
	errorEmpty := errors.New("empty phone number input")

	replacer := strings.NewReplacer(":", "", ".", "", " ", "", "-", "", "+", "", "(", "", ")", "")
	sanitizedPhoneNumber := replacer.Replace(phoneNumber)

	if len(sanitizedPhoneNumber) == 0 {
		return "", errorEmpty
	}

	if len(sanitizedPhoneNumber) > 10 && sanitizedPhoneNumber[:1] == "1" {
		sanitizedPhoneNumber = sanitizedPhoneNumber[1:]
	}

	if len(sanitizedPhoneNumber) > 10 {
		return sanitizedPhoneNumber, errorLength
	}

	if sanitizedPhoneNumber[:1] == "1" || sanitizedPhoneNumber[:1] == "0" {
		return sanitizedPhoneNumber, errorAreaCode
	}

	if sanitizedPhoneNumber[3:4] == "1" || sanitizedPhoneNumber[3:4] == "0" {
		return sanitizedPhoneNumber, errorExchangeCode
	}

	return sanitizedPhoneNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)

	if err != nil {
		return phoneNumber, err
	}

	return phoneNumber[:3], nil
}

func Format(phoneNumber string) (string, error) {

	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		fmt.Println(err)
	}
	areaCode, err := AreaCode(phoneNumber)
	if err != nil {
		fmt.Println(err)
	}
	exchangeCode := phoneNumber[3:6]
	subscriberID := phoneNumber[6:10]

	return fmt.Sprintf("(%s) %s-%s", areaCode, exchangeCode, subscriberID), err
}
