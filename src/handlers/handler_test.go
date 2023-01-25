package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kalpit-sharma-dev/parkinglot-service/src/mocks"
	"github.com/kalpit-sharma-dev/parkinglot-service/src/models"

	"github.com/golang/mock/gomock"
)

func TestHandleCreateSlotEvent(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockParkingLotService(ctrl)
	urlHandler := NewHandler(mockService)

	var successReq = models.Slot{
		SlotSize: 10,
	}

	tests := []struct {
		name           string
		req            models.Slot
		setupMocks     func()
		expectedStatus int
	}{
		{
			name: "Success",
			req:  successReq,
			setupMocks: func() {
				mockService.EXPECT().CreateSlotEvent(gomock.Any(), gomock.Any()).Return(models.Slot{}, nil).AnyTimes()
			},
			expectedStatus: 201,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()
			b, _ := json.Marshal(tt.req)
			rr := httptest.NewRecorder()
			httpReq, _ := http.NewRequest("POST", "/url", bytes.NewBuffer(b))
			urlHandler.HandleCreateSlotEvent(rr, httpReq)
			if tt.expectedStatus != rr.Code {
				t.Errorf("Error got = %v and want = %v", rr.Code, tt.expectedStatus)
			}
		})
	}
}
