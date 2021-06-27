package usecase

import (
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/myusufirfanh/go-demo-service/domain/mocks/mock_addon"
	"github.com/myusufirfanh/go-demo-service/domain/mocks/mock_product"
	"github.com/myusufirfanh/go-demo-service/models"
	"github.com/myusufirfanh/go-demo-service/shared/config"
	"github.com/myusufirfanh/go-demo-service/shared/container"
	"github.com/myusufirfanh/go-demo-service/shared/util"
)

// Testing unexported (private) function
func TestCalculateTax(t *testing.T) {
	expectedResult := int64(5)
	actualResult := calculateTax(50)
	if actualResult != expectedResult {
		t.Errorf("Expected different than actual result! Expected: %d - Actual: %d", expectedResult, actualResult)
	}
}

var conf config.ImmutableConfigInterface

// TestMain function will be called once, BEFORE running all tests
func TestMain(m *testing.M) {
	// This part is called setup
	container := container.DefaultContainer()
	conf = container.MustGet("shared.config").(config.ImmutableConfigInterface)

	// m.Run() will run the actual tests
	exitVal := m.Run()

	// Any code here is called teardown
	// Put code here

	os.Exit(exitVal)
}

// Testing a method
func TestGetAddonsByProductID(t *testing.T) {

	// Gomock
	// Must be init'd in each test because it needs t parameter
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_product.NewMockRepository(mockCtrl)
	mockAddon := mock_addon.NewMockUsecase(mockCtrl)

	// Supply those mocks to our usecase
	ucase := usecase{
		repository:   mockRepo,
		addonUsecase: mockAddon,
	}

	// Construct (custom) context
	c := &util.CustomApplicationContext{
		UserJWT:      &models.UserJWT{ID: 1},
		MysqlSession: nil,
		SharedConf:   conf,
	}

	// Initiate mock data from JSON
	addonsResult := []models.Addon{}
	err := util.ReadJSON("../../../files/test/addons_data.json", &addonsResult)

	const productID = int64(1)
	const expectProductName = "Cheeseburger"

	// Expect the behavior of mocked function
	mockAddon.EXPECT().GetAllAddonsByProductID(gomock.Any(), gomock.Eq(productID)).Return(addonsResult, nil).Times(1)

	// Call the function we want to test
	actualResult, err := ucase.GetAddonsByProductID(c, productID)
	if len(actualResult) != 3 {
		t.Error("Expected different than actual result!")
	}
	if err != nil {
		t.Error("Expected different than actual result!")
	}
}

// Testing with sqlmock
func TestGetProduct(t *testing.T) {

	// Gomock
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_product.NewMockRepository(mockCtrl)
	mockAddon := mock_addon.NewMockUsecase(mockCtrl)

	// Supply those mocks to our usecase
	ucase := usecase{
		repository:   mockRepo,
		addonUsecase: mockAddon,
	}

	// SQL mock
	db, sqlmock, err := sqlmock.New()
	gdb, err := gorm.Open("mysql", db)

	sqlmock.ExpectBegin()
	sqlmock.ExpectCommit()

	// Construct (custom) context, supply MySQLSession from sqlmock
	c := &util.CustomApplicationContext{
		UserJWT:      &models.UserJWT{ID: 1},
		MysqlSession: gdb,
		SharedConf:   conf,
	}

	// Initiate mock data from JSON
	productResult := &models.Product{}
	err = util.ReadJSON("../../../files/test/product_data.json", &productResult)

	const productID = int64(1)
	const expectProductName = "Cheeseburger"

	// Expect the behavior of mocked function
	mockRepo.EXPECT().GetProduct(gomock.Any(), gomock.Eq(productID)).Return(productResult, nil).AnyTimes()

	// Call the function we want to test
	actualResult, err := ucase.GetProductWithTx(c, productID)
	if actualResult.ID != productID {
		t.Error("Expected different than actual result!")
	}
	if actualResult.ProductName != expectProductName {
		t.Error("Expected different than actual result!")
	}
	if err != nil {
		t.Error("Expected different than actual result!")
	}
}
