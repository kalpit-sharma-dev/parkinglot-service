package repository

import (
	"context"

	"github.com/kalpit-sharma-dev/parkinglot-service/src/models"
)

type DatabaseRepository interface {
	CreateSlotEvent(ctx context.Context, req models.Slot) (vehicle models.Slot, err error)
	GetAllCarsWithColor(ctx context.Context, reqColor string) (vehicle []models.Vehicle, err error)
	GetSlotNumberWithCarID(ctx context.Context, reqNumber string) (vehicle models.Vehicle, err error)
	GetAllSlotNumberWithColor(ctx context.Context, reqColor string) (vehicle []models.Vehicle, err error)

	CreateParkEvent(ctx context.Context, req models.Vehicle) (vehicle models.Vehicle, err error)
	ExitParkEvent(ctx context.Context, req models.Vehicle) (vehicle models.Vehicle, err error)
}
