package errors

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleError(t *testing.T) {
	type args struct {
		scenario           string
		inputError         error
		expectedResponse   error
		expectedStatusCode int
	}

	tests := []args{

		{
			scenario:   "Internal Server Error",
			inputError: errors.New("Internal Server Error"),
			//expectedResponse:   http.StatusInternalServerError,
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {

			res := httptest.NewRecorder()
			HandleError(res, test.inputError)

		})
	}
}
