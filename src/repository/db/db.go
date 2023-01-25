package db

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/kalpit-sharma-dev/parkinglot-service/src/constants"
	"github.com/kalpit-sharma-dev/parkinglot-service/src/models"
	"github.com/kalpit-sharma-dev/parkinglot-service/src/repository"
)

//var Atom atomic.Int32

var SlotSize int

var PeakHours map[string]int

type Vehicles struct {
	Color     string
	Number    string
	Slot      string
	CreatedAt time.Time
}

var SlotDB map[string]Vehicles
var EmptySlotDB map[string]string

func init() {

	SlotDB = make(map[string]Vehicles)
	EmptySlotDB = make(map[string]string)

}

type DatabaseImpl struct {
	DBObj *[]Vehicles
	RW    sync.RWMutex
	Day   map[int]int //key value pair for day and earning
	Month map[int]int ////key value pair for day and earning
}

func (db *DatabaseImpl) CreateSlotEvent(ctx context.Context, req models.Slot) (vehicle models.Slot, err error) {

	db.RW.Lock()
	oldSlotSize := SlotSize

	SlotSize = req.SlotSize + SlotSize

	vehicle.SlotSize = SlotSize
	for i := oldSlotSize + 1; i <= SlotSize; i++ {
		slot := strconv.Itoa(i)
		EmptySlotDB[slot] = slot
	}

	db.RW.Unlock()

	return
}

// CreateEvents implements repository.DatabaseRepository
func (db *DatabaseImpl) GetAllCarsWithColor(ctx context.Context, reqColor string) (vehicle []models.Vehicle, err error) {

	db.RW.RLock()
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
	db.RW.RUnlock()
	return
}

func (db *DatabaseImpl) GetSlotNumberWithCarID(ctx context.Context, reqNumber string) (vehicle models.Vehicle, err error) {
	db.RW.RLock()
	for _, value := range *db.DBObj {
		if value.Number == reqNumber {
			vehicle = models.Vehicle{
				Slot:   value.Slot,
				Number: value.Number,
				Color:  value.Color,
				//CreatedAt: value.CreatedAt,
			}
			db.RW.RUnlock()
			return
		}
	}

	if len(strings.TrimSpace(vehicle.Slot)) == 0 {
		db.RW.RUnlock()
		return models.Vehicle{}, errors.New("NO Vehicle Found for this ID")
	}
	db.RW.RUnlock()
	return
}

// GetAllEvents implements repository.DatabaseRepository
func (db *DatabaseImpl) GetAllSlotNumberWithColor(ctx context.Context, reqColor string) (vehicle []models.Vehicle, err error) {

	db.RW.RLock()
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
	db.RW.RUnlock()
	return

}

func (db *DatabaseImpl) CreateParkEvent(ctx context.Context, req models.Vehicle) (vehicle models.Vehicle, err error) {

	db.RW.Lock()
	latestEmptySlot, err := fetchEmptySlot()

	if err != nil {
		db.RW.Unlock()
		return vehicle, err

	}
	value := Vehicles{
		Color:     req.Color,
		Number:    req.Number,
		Slot:      latestEmptySlot,
		CreatedAt: time.Now(),
	}

	*db.DBObj = append(*db.DBObj, value)
	SlotDB[latestEmptySlot] = value
	delete(EmptySlotDB, latestEmptySlot)
	db.RW.Unlock()
	vehicle.Color = value.Color
	vehicle.Number = value.Number
	vehicle.Slot = value.Slot
	return
}

func (db *DatabaseImpl) ExitParkEvent(ctx context.Context, req models.Vehicle) (models.Vehicle, error) {

	var vehicle models.Vehicle
	db.RW.Lock()
	for index, value := range *db.DBObj {
		if value.Number == req.Number {
			delete(SlotDB, req.Slot)
			addEmptySlot(req.Slot)
			db.DBObj = remove(db.DBObj, index)
			fmt.Println("###################", index, value)
			vehicle.Color = value.Color
			vehicle.Number = value.Number
			vehicle.Slot = value.Slot
			db.RW.Unlock()
			return vehicle, nil
		}
	}
	db.RW.Unlock()
	err := errors.New(constants.NoCarParkedError)
	return vehicle, err
}

// NewRepository returns instance of Database repository
func NewDatabaseRepository(db *[]Vehicles) repository.DatabaseRepository {

	return &DatabaseImpl{DBObj: db}
}

func remove(slice *[]Vehicles, slot int) *[]Vehicles {
	(*slice) = append((*slice)[:slot], (*slice)[slot+1:]...)
	return slice
}

func fetchEmptySlot() (string, error) {

	for k, _ := range EmptySlotDB {
		delete(EmptySlotDB, k)
		return k, nil
	}
	return "", errors.New("no empty slot found")
}

func addEmptySlot(inputSlot string) {

	EmptySlotDB[inputSlot] = inputSlot
}
