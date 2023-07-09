package helpers

import (
	"reflect"
	"time"

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

func ToDateTime(datetime string) time.Time {
	dt, err := time.Parse("2006-01-02 15:04:05", datetime)

	if err != nil {
		panic("could not parse datetime")
	}

	return dt
}

func ToFormattedDateTime(datetime time.Time) string {
	return datetime.Format("2006-01-02 15:04:05")
}
