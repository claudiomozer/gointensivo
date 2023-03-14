package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnAnErrorIfIdIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "invalid id")
}

func TestShouldReturnAnErrorIfPriceIsBlank(t *testing.T) {
	order := Order{ID: "teste"}
	assert.Error(t, order.Validate(), "invalid price")
}

func TestShouldReturnAnErrorIfTaxIsBlank(t *testing.T) {
	order := Order{ID: "teste", Price: 13.9}
	assert.Error(t, order.Validate(), "invalid tax")
}

func TestShouldCreateAnOrderOnSuccess(t *testing.T) {
	order := Order{ID: "teste", Price: 13.9, Tax: 1.0}
	order.CalculateFinalPrice()
	assert.Equal(t, order.FinalPrice, 14.9)
}
