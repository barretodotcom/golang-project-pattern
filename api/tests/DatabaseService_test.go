package tests

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang-project-pattern/api/mocks"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateStudentOk(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	service := mocks.NewMockIDatabase(mockCtrl)
	filter := bson.D{{Key: "name", Value: "Lucas Barreto"}}

	service.EXPECT().
		CountStudents(filter).
		Return(int64(1)).
		Times(1)

	countResult := int64(1)

	actualOutput := service.CountStudents(filter)

	assert.Equal(t, countResult, actualOutput)

}
