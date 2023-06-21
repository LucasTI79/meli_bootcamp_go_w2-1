package helpers

import "reflect"

func Fill(first interface{}, second interface{}) interface{} {
	valueOfFirst := reflect.ValueOf(first)
	if valueOfFirst.IsNil() {
		return second
	}

	return valueOfFirst.Elem().Interface()
}
