package helpers

import (
	"reflect"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Fill(first interface{}, second interface{}) interface{} {
	valueOfFirst := reflect.ValueOf(first)
	if valueOfFirst.IsNil() {
		return second
	}

	return valueOfFirst.Elem().Interface()
}

func ToFormattedAddress(address string) string {
	caser := cases.Title(language.BrazilianPortuguese)
	return caser.String(address)
}
