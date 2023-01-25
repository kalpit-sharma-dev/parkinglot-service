package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/kalpit-sharma-dev/parkinglot-service/src/mocks"
	"github.com/kalpit-sharma-dev/parkinglot-service/src/models"

	"github.com/golang/mock/gomock"
)

func TestCreateSlot(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(mockCtrl)
	VehicleService := NewParkingLotService(mockRepo)

	type args struct {
		ctx        context.Context
		VehicleReq models.Slot
	}
	tests := []struct {
		name        string
		args        args
		setupMocks  func()
		wantUrlResp models.Slot
		wantErr     bool
	}{

		{
			name: "Failure",
			args: args{
				ctx: context.Background(),
				VehicleReq: models.Slot{
					SlotSize: 0,
				},
			},
			setupMocks: func() {
				mockRepo.EXPECT().CreateSlotEvent(gomock.Any(), gomock.Any()).Return(models.Slot{}, errors.New("Some error")).Times(1)

			},
			wantUrlResp: models.Slot{},
			wantErr:     true,
		},
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				VehicleReq: models.Slot{
					SlotSize: 1,
				},
			},
			setupMocks: func() {
				mockRepo.EXPECT().CreateSlotEvent(gomock.Any(), gomock.Any()).Return(models.Slot{
					SlotSize: 1,
				}, nil).Times(1)

			},
			wantUrlResp: models.Slot{
				SlotSize: 1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()
			gotUrlResp, err := VehicleService.CreateSlotEvent(tt.args.ctx, tt.args.VehicleReq)
			if (err != nil) != tt.wantErr {
				if err != errors.New("Some error") {
					t.Errorf("CreateVehicle() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(gotUrlResp, tt.wantUrlResp) {
				t.Errorf("CreateVehicle() got = %v, want %v", gotUrlResp, tt.wantUrlResp)
			}

		})
	}
}

func TestGetAllVehicles(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(mockCtrl)
	VehicleService := NewParkingLotService(mockRepo)

	type args struct {
		ctx        context.Context
		VehicleReq models.Vehicle
	}
	tests := []struct {
		name        string
		args        args
		setupMocks  func()
		wantUrlResp []models.Vehicle
		wantErr     bool
	}{

		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				VehicleReq: models.Vehicle{
					Slot:   "1",
					Number: "123",
					Color:  "Black",
				},
			},
			setupMocks: func() {
				mockRepo.EXPECT().GetAllSlotNumberWithColor(gomock.Any(), gomock.Any()).Return([]models.Vehicle{{Slot: "1",
					Number: "123",
					Color:  "Black"}}, nil).Times(1)

			},
			wantUrlResp: []models.Vehicle{
				{
					Slot:   "1",
					Number: "123",
					Color:  "Black",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()
			gotUrlResp, err := VehicleService.GetAllSlotNumberWithColor(tt.args.ctx, "Black")
			if (err != nil) != tt.wantErr {
				if err != errors.New("Some error") {
					t.Errorf("GetAllVehicles() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(gotUrlResp, tt.wantUrlResp) {
				t.Errorf("GetAllVehicles() got = %v, want %v", gotUrlResp, tt.wantUrlResp)
			}

		})
	}
}
