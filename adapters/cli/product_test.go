package cli_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/kalleocarrilho/go-hexagonal/adapters/cli"
	mock_application "github.com/kalleocarrilho/go-hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)

	productName := "Product Test"
	productPrice := 25.99
	productStatus := "enabled"
	productId := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)

	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()

	serviceMock := mock_application.NewMockProductServiceInterface(ctrl)

	serviceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	expectedResult := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s.",
		productId, productName, productPrice, productStatus)

	result, err := cli.Run(serviceMock, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf("Product %s has been enabled.", productName)

	result, err = cli.Run(serviceMock, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf("Product %s has been disabled.", productName)

	result, err = cli.Run(serviceMock, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		productId, productName, productPrice, productStatus)

	result, err = cli.Run(serviceMock, "", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

}
