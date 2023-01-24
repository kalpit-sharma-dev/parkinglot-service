package db

import (
	"context"
	"errors"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/kalpit-sharma-dev/parkinglot-service/src/models"
	"github.com/kalpit-sharma-dev/parkinglot-service/src/repository"
)

//var Atom atomic.Int32

var SlotSize chan int

type Vehicles struct {
	Color     string
	Number    string
	Slot      string
	CreatedAt time.Time
}

type Color map[string][]string

type Car map[string]Color
type DatabaseImpl struct {
	DBObj *[]Vehicles
	RW    sync.RWMutex
	Day   map[int]int //key value pair for day and earning
	Month map[int]int ////key value pair for day and earning
}

// CreateEvents implements repository.DatabaseRepository
func (db *DatabaseImpl) GetAllCarsWithColor(ctx context.Context, reqColor string) (vehicle []models.Vehicle, err error) {

	//db.RW.RLock()
	for _, value := range *db.DBObj {
		if value.Color == reqColor {
			vehicle = append(vehicle, models.Vehicle{
				Slot:   value.Slot,
				Number: value.Number,
				Color:  value.Color,
				// /CreatedAt: value.CreatedAt,
			})
		}
	}
	//db.RW.RUnlock()
	return
}

func (db *DatabaseImpl) GetSlotNumberWithCarID(ctx context.Context, reqNumber string) (vehicle models.Vehicle, err error) {
	//db.RW.RLock()
	for _, value := range *db.DBObj {
		if value.Number == reqNumber {
			vehicle = models.Vehicle{
				Slot:   value.Slot,
				Number: value.Number,
				Color:  value.Color,
				//CreatedAt: value.CreatedAt,
			}
			//db.RW.RUnlock()
			return
		}
	}

	if len(strings.TrimSpace(vehicle.Slot)) == 0 {
		return models.Vehicle{}, errors.New("NO Vehicle Found for this ID")
	}
	return
}

// GetAllEvents implements repository.DatabaseRepository
func (db *DatabaseImpl) GetAllSlotNumberWithColor(ctx context.Context, reqColor string) (vehicle []models.Vehicle, err error) {

	//db.RW.RLock()
	for _, value := range *db.DBObj {
		if value.Color == reqColor {
			vehicle = append(vehicle, models.Vehicle{
				Slot:   value.Slot,
				Number: value.Number,
				Color:  value.Color,
				//CreatedAt: value.CreatedAt,
			})
		}
	}
	//db.RW.RUnlock()
	return

}

func (db *DatabaseImpl) CreateParkEvent(ctx context.Context, req models.Vehicle) (vehicle models.Vehicle, err error) {

	db.RW.Lock()
	time.Sleep(4 * time.Second)
	value := Vehicles{
		Color:     req.Color,
		Number:    req.Number,
		Slot:      req.Slot,
		CreatedAt: time.Now(),
	}
	*db.DBObj = append(*db.DBObj, value)
	db.RW.Unlock()
	vehicle.Color = value.Color
	vehicle.Number = value.Number
	vehicle.Slot = value.Slot
	return
}

func (db *DatabaseImpl) ExitParkEvent(ctx context.Context, req models.Vehicle) (vehicle models.Vehicle, err error) {

	db.RW.Lock()
	for _, value := range *db.DBObj {
		if value.Number == req.Number {
			slot, err := strconv.Atoi(value.Slot)
			if err != nil {
				log.Println("error exit for car number", req.Number, err)
				return vehicle, err
			}
			remove(db.DBObj, slot-1)
			db.RW.Unlock()
			return vehicle, err
		}
	}
	db.RW.Unlock()

	return
}

// NewRepository returns instance of Database repository
func NewDatabaseRepository(db *[]Vehicles) repository.DatabaseRepository {

	return &DatabaseImpl{DBObj: db}
}

func remove(slice *[]Vehicles, slot int) []Vehicles {
	return append((*slice)[:slot], (*slice)[slot+1:]...)
}
