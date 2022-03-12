package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("40000000000", "Jose da Silva", 12, 2024, 123)

	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("4909991766000586", "Jose da Silva", 12, 2024, 123)

	assert.Nil(t, err)

}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("4909991766000586", "Jose da Silva", 13, 2024, 123)

	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("4909991766000586", "Jose da Silva", 0, 2024, 123)

	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("4909991766000586", "Jose da Silva", 6, 2024, 123)

	assert.Nil(t, err)

}

func TestCreditCardExpirationYear(t *testing.T) {

	lastYear := time.Now().Year() - 1

	_, err := NewCreditCard("4909991766000586", "Jose da Silva", 12, lastYear, 123)

	assert.Equal(t, "invalid expiration year", err.Error())

	_, err = NewCreditCard("4909991766000586", "Jose da Silva", 6, 2024, 123)

	assert.Nil(t, err)

}
