package usecase_test

import (
	"encoding/json"
	"testing"

	"github.com/myusufirfanh/go-demo-service/domain/product/usecase"
	"github.com/myusufirfanh/go-demo-service/models"
	"github.com/nsf/jsondiff"
	"github.com/stretchr/testify/assert"
)

func TestCalculateDeliveryFeeWithAssert(t *testing.T) {
	expectedResult := int64(10)
	actualResult := usecase.CalculateDeliveryFee(50)
	assert.Equal(t, expectedResult, actualResult)
}

func TestJSON(t *testing.T) {
	input := models.Product{
		ID:          1,
		ProductName: "Ice Cream",
		Description: "Vanilla Flavor",
		Price:       500,
	}
	expectedJsonStr := `{
			"price": 1500,
			"id": 1,
			"product_name": "Ice Cream",
			"description": "Vanilla Flavor"
		}`
	out := usecase.AlterProductPrice(input)
	outJsonStr, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		t.Fatal("error marshaling package", err)
	}

	diffOpts := jsondiff.DefaultConsoleOptions()
	res, diff := jsondiff.Compare([]byte(expectedJsonStr), []byte(outJsonStr), &diffOpts)

	if res != jsondiff.FullMatch {
		t.Errorf("the expected result is not equal to what we have: %s", diff)
	}

}
