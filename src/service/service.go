package service

import (
	"context"
	"errors"
	"strings"

	"github.com/kalpit-sharma-dev/parkinglot-service/src/models"
	"github.com/kalpit-sharma-dev/parkinglot-service/src/repository"
)

type ParkingLotServiceImpl struct {
	repository repository.DatabaseRepository
}

func (s *ParkingLotServiceImpl) CreateSlotEvent(ctx context.Context, req models.Slot) (vehicle models.Slot, err error) {
	vehicle, err = s.repository.CreateSlotEvent(ctx, req)
	return
}

// CreateParkingLot implements ParkingLotService
func (s *ParkingLotServiceImpl) GetAllCarsWithColor(ctx context.Context, reqColor string) (vehicle []models.Vehicle, err error) {
	vehicle, err = s.repository.GetAllCarsWithColor(ctx, reqColor)
	return
}

func (s *ParkingLotServiceImpl) GetSlotNumberWithCarID(ctx context.Context, reqNumber string) (vehicle models.Vehicle, err error) {
	vehicle, err = s.repository.GetSlotNumberWithCarID(ctx, reqNumber)
	return
}

// GetParkingLot implements ParkingLotService
func (s *ParkingLotServiceImpl) GetAllSlotNumberWithColor(ctx context.Context, reqColor string) (vehicle []models.Vehicle, err error) {
	if len(strings.TrimSpace(reqColor)) == 0 {
		return vehicle, errors.New("empty color recieved in query")
	}
	vehicle, err = s.repository.GetAllSlotNumberWithColor(ctx, reqColor)
	return
}

func (s *ParkingLotServiceImpl) CreateParkEvent(ctx context.Context, req models.Vehicle) (vehicle models.Vehicle, err error) {
	vehicle, err = s.repository.CreateParkEvent(ctx, req)
	return
}

func (s *ParkingLotServiceImpl) ExitParkEvent(ctx context.Context, req models.Vehicle) (vehicle models.Vehicle, err error) {
	vehicle, err = s.repository.ExitParkEvent(ctx, req)
	return
}

func NewParkingLotService(repository repository.DatabaseRepository) ParkingLotService {
	return &ParkingLotServiceImpl{repository: repository}
}
