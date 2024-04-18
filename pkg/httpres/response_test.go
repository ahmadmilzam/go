package httpres

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type sampleData struct {
	ID    string
	Name  string
	Phone string
}

type dataArr []sampleData

func TestGenerateOK(t *testing.T) {
	data := sampleData{
		ID:    "123",
		Name:  "John Doe",
		Phone: "+6281281281200",
	}

	data2 := sampleData{
		ID:    "456",
		Name:  "Jane Doe",
		Phone: "+6281281281211",
	}

	dataSlice := dataArr{data, data2}

	actualFlat := GenerateOK(data)
	actualArr := GenerateOK(dataSlice)

	expectedFlat := HttpResponse{
		Success: true,
		Data:    data,
	}
	expectedArr := HttpResponse{
		Success: true,
		Data:    dataSlice,
	}
	// fe.InitServiceCode(issuer.TransactionManagementService)
	assert.Equal(t, expectedFlat, actualFlat)
	assert.Equal(t, expectedArr, actualArr)
}

func TestGenerateErr(t *testing.T) {
	errMessage := "Bad Request"
	err := fmt.Errorf("%s_%w", GenericBadRequest, errors.New("BAD_REQUEST"))
	actual := GenerateErrResponse(err, errMessage)
	expected := HttpResponse{
		Success: false,
		Error: &ErrorDetails{
			Code:    GetCaseCode(err),
			Message: errMessage,
		},
	}
	// fe.InitServiceCode(issuer.TransactionManagementService)
	assert.Equal(t, expected, actual)
}

func TestGetStatusCode(t *testing.T) {
	assert.Equal(t, http.StatusBadRequest, GetStatusCode(errors.New(GenericBadRequest)))
	assert.Equal(t, http.StatusUnauthorized, GetStatusCode(errors.New(GenericUnauthorized)))
	assert.Equal(t, http.StatusNotFound, GetStatusCode(errors.New(GenericNotFound)))
	assert.Equal(t, http.StatusMethodNotAllowed, GetStatusCode(errors.New(GenericMethodNotAllowed)))
	assert.Equal(t, http.StatusConflict, GetStatusCode(errors.New(GenericConflict)))
	assert.Equal(t, http.StatusUnprocessableEntity, GetStatusCode(errors.New(GenericUnprocessable)))
	assert.Equal(t, http.StatusTooManyRequests, GetStatusCode(errors.New(GenericTooManyRequests)))
	assert.Equal(t, http.StatusInternalServerError, GetStatusCode(errors.New(GenericInternalError)))
}
