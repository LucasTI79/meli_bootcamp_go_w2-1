package helpers_test

import (
	"testing"
	"time"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/helpers"
	"github.com/stretchr/testify/assert"
)

func TestFill(t *testing.T) {
	t.Run("Should return first value when first is not nil", func(t *testing.T) {
		first := "first"
		firstPtr := &first
		second := "second"

		result := helpers.Fill(firstPtr, second)

		assert.Equal(t, first, result)
	})

	t.Run("Should return second value when first is nil", func(t *testing.T) {
		var firstPtr *string
		second := "second"

		result := helpers.Fill(firstPtr, second)

		assert.Equal(t, second, result)
	})
}

func TestToFormattedAddress(t *testing.T) {
	t.Run("Should return the formatted address", func(t *testing.T) {
		unformattedAddress := "eXaMple address"
		expectedResult := "Example Address"

		result := helpers.ToFormattedAddress(unformattedAddress)

		assert.Equal(t, expectedResult, result)
	})
}

func TestToDateTime(t *testing.T) {
	t.Run("Should return a datetime object based on the provided string datetime format yyyy-mm-dd HH:mm:ss", func(t *testing.T) {
		datetimeString := "2021-01-01 01:00:00"
		expectedResult := time.Date(2021, 1, 1, 1, 0, 0, 0, time.UTC)

		result := helpers.ToDateTime(datetimeString)

		assert.Equal(t, expectedResult, result)
	})

	t.Run("Should panic when the provided string datetime is in wrong format", func(t *testing.T) {
		datetimeString := "2021-01-01 01:00"

		assert.Panics(t, func() { helpers.ToDateTime(datetimeString) })
	})
}

func TestToFormattedDateTime(t *testing.T) {
	t.Run("Should return the date in yyyy-mm-dd HH:mm:ss format based on the provied datetime object", func(t *testing.T) {
		datetime := time.Date(2021, 1, 1, 1, 0, 0, 0, time.UTC)
		expectedResult := "2021-01-01 01:00:00"

		result := helpers.ToFormattedDateTime(datetime)

		assert.Equal(t, expectedResult, result)
	})
}
