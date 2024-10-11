package application_test

import (
	"testing"

	uuid "github.com/google/uuid"
	"github.com/kalleocarrilho/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.New().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "invalid status"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = float64(-30)
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())

	product.Price = 10
	_, err = product.IsValid()
	require.Nil(t, err)

	product.ID = "Some ID"
	valid, err := product.IsValid()
	require.Equal(t, false, valid)
}

func TestProduct_Getters(t *testing.T) {
	id := uuid.New().String()
	product := application.Product{}
	product.ID = id
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	require.Equal(t, id, product.GetID())
	require.Equal(t, "Hello", product.GetName())
	require.Equal(t, application.DISABLED, product.GetStatus())
	require.Equal(t, float64(10), product.GetPrice())
}
