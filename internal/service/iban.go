package service

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

type IbanStructure struct {
	lengthMap map[string]int
	letters   map[string]string
}

var ibanLenght IbanStructure

func InitService() {
	ibanLenght.lengthMap = make(map[string]int)
	ibanLenght.lengthMap["TR"] = 24
	ibanLenght.lengthMap["DE"] = 22

	ibanLenght.letters = make(map[string]string)
	ibanLenght.letters["A"] = "10"
	ibanLenght.letters["B"] = "11"
	ibanLenght.letters["C"] = "12"
	ibanLenght.letters["D"] = "13"
	ibanLenght.letters["E"] = "14"
	ibanLenght.letters["F"] = "15"
	ibanLenght.letters["G"] = "16"
	ibanLenght.letters["H"] = "17"
	ibanLenght.letters["I"] = "18"
	ibanLenght.letters["J"] = "19"
	ibanLenght.letters["K"] = "20"
	ibanLenght.letters["L"] = "21"
	ibanLenght.letters["M"] = "22"
	ibanLenght.letters["N"] = "23"
	ibanLenght.letters["O"] = "24"
	ibanLenght.letters["P"] = "25"
	ibanLenght.letters["Q"] = "26"
	ibanLenght.letters["R"] = "27"
	ibanLenght.letters["S"] = "28"
	ibanLenght.letters["T"] = "29"
	ibanLenght.letters["U"] = "30"
	ibanLenght.letters["V"] = "31"
	ibanLenght.letters["W"] = "32"
	ibanLenght.letters["X"] = "33"
	ibanLenght.letters["Y"] = "34"
	ibanLenght.letters["Z"] = "35"
}

func ValidateIban(iban string) (bool, string) {

	iban = strings.ReplaceAll(iban, " ", "")

	country := iban[0:2]
	country = strings.ToUpper(country)

	if !isLetter(country) {
		return true, "Iban should start with 2 letters"
	}

	ibanCountryLength := ibanLenght.lengthMap[country]
	if ibanCountryLength == 0 {
		return true, "Country not supported:" + country
	}
	if ibanCountryLength != len(iban) {
		return true, "Iban should be with length " + strconv.Itoa(ibanCountryLength) + " for country" + country
	}

	ibanToBeValidated := iban[4:] + iban[0:4]
	firstLetterCountry := ibanLenght.letters[string(country[0])]
	secondCountry := ibanLenght.letters[string(country[1])]

	ibanToBeValidated = ibanToBeValidated[0:len(ibanToBeValidated)-4] + firstLetterCountry + secondCountry + ibanToBeValidated[len(ibanToBeValidated)-2:]
	ibanFirstInt, _ := strconv.ParseInt(ibanToBeValidated[0:10], 10, 64)

	pow := int64(math.Pow(10, float64(len(ibanToBeValidated)-10)))
	ibanFirstInt = ibanFirstInt%97
	pow = pow%97
	firstHalf := (ibanFirstInt * pow) % 97

	ibanSecondInt, _ := strconv.ParseInt(ibanToBeValidated[10:], 10, 64)
	ibanSecondInt = ibanSecondInt %97

	result := (firstHalf + ibanSecondInt)%97

	if result != 1 {
		return true, "Iban is not valid"
	}

	return false, ""

}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
