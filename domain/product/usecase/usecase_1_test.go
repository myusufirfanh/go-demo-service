package usecase_test

import (
	"testing"

	"github.com/myusufirfanh/go-demo-service/domain/product/usecase"
)

// Testing a normal exported function
func TestCalculateDeliveryFee(t *testing.T) {
	expectedResult := int64(10)
	actualResult := usecase.CalculateDeliveryFee(50)
	if actualResult != expectedResult {
		t.Errorf("Expected different than actual result! Expected: %d - Actual: %d", expectedResult, actualResult)
	}
}

// Testing with "test table"
func TestCalculateDeliveryFeeWithTable(t *testing.T) {
	tables := []struct {
		input  int64
		output int64
	}{
		{50, 10},
		{100, 20},
		{300, 60},
	}

	for _, table := range tables {
		actualResult := usecase.CalculateDeliveryFee(table.input)
		if actualResult != table.output {
			t.Errorf("Expected different than actual result! Expected: %d - Actual: %d", table.output, actualResult)
		}
	}
}

// Testing with subtests
func TestCalculateDeliveryFeeWithSubtests(t *testing.T) {

	t.Run("Input 50", func(t *testing.T) {
		expectedResult := int64(10)
		actualResult := usecase.CalculateDeliveryFee(50)
		if actualResult != expectedResult {
			t.Errorf("Expected different than actual result! Expected: %d - Actual: %d", expectedResult, actualResult)
		}
	})

	t.Run("Input 100", func(t *testing.T) {
		expectedResult := int64(20)
		actualResult := usecase.CalculateDeliveryFee(100)
		if actualResult != expectedResult {
			t.Errorf("Expected different than actual result! Expected: %d - Actual: %d", expectedResult, actualResult)
		}
	})

}
