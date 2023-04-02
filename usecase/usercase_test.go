package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	publicHandler "nexter-assignment/handler/http/public"
	"nexter-assignment/tests/mocks"
	"nexter-assignment/usecase"
)

// TestUsecaseSuite ..
type TestUsecaseSuite struct {
	suite.Suite
	ctrl *gomock.Controller

	mockRepo *mocks.MockRepository

	publicUC publicHandler.Usecase
}

// TestUsecase ..
func TestUsecase(t *testing.T) {
	suite.Run(t, new(TestUsecaseSuite))
}

// BeforeTest ..
func (t *TestUsecaseSuite) BeforeTest(suiteName string, testName string) {
	t.ctrl = gomock.NewController(t.T())
	t.mockRepo = mocks.NewMockRepository(t.ctrl)
	t.publicUC = usecase.NewUsecase(t.mockRepo)
}

// AfterTest ..
func (t *TestUsecaseSuite) AfterTest(suiteName string, testName string) {
	t.ctrl.Finish()
}
