package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
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
	Model     string
	CreatedAt time.Time
}

var SlotDB map[string]Vehicles
var EmptySlotDB map[string]string

func init() {

	SlotDB = make(map[string]Vehicles)
	EmptySlotDB = make(map[string]string)

}

type DatabaseImpl struct {
	DBObj      *[]Vehicles
	DBmySqlCon *sql.DB
	RW         sync.RWMutex
	Day        map[int]int //key value pair for day and earning
	Month      map[int]int ////key value pair for day and earning
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
				Model:  value.Model,
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
				Model:  value.Model,
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
				Model:  value.Model,
				//CreatedAt: value.CreatedAt,
			})
		}
	}
	db.RW.RUnlock()

	rows, err := db.DBmySqlCon.Query("SELECT car_number,color,model FROM vehicles WHERE color = ?", reqColor)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}
	defer rows.Close()

	// Process the results
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&vehicle); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	return

}

func (db *DatabaseImpl) CreateParkEvent(ctx context.Context, req models.Vehicle) (vehicle models.Vehicle, err error) {

	fmt.Println("CreateParkEvent  |  Insert user into the database")
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
		Model:     req.Model,
	}

	*db.DBObj = append(*db.DBObj, value)
	SlotDB[latestEmptySlot] = value
	delete(EmptySlotDB, latestEmptySlot)
	db.RW.Unlock()
	vehicle.Color = value.Color
	vehicle.Number = value.Number
	vehicle.Slot = value.Slot
	vehicle.Model = value.Model
	// Insert slots into the database
	fmt.Println(" Insert slots into the database")
	_, err = db.DBmySqlCon.Exec("INSERT INTO slots (slot_number,car_number) VALUES (?,?)", value.Slot, value.Number)
	if err != nil {
		log.Println(err)
		return
	}
	// Insert vehicles into the database
	fmt.Println(" Insert vehicles into the database")
	_, err = db.DBmySqlCon.Exec("INSERT INTO vehicles (car_number,color,model) VALUES (?,?,?)", value.Number, value.Color, value.Model)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (db *DatabaseImpl) ExitParkEvent(ctx context.Context, req models.Vehicle) (models.Vehicle, error) {
	flag := false
	var vehicle models.Vehicle
	db.RW.Lock()
	for index, value := range *db.DBObj {
		if value.Number == req.Number {
			flag = true
			delete(SlotDB, req.Slot)
			addEmptySlot(req.Slot)
			db.DBObj = remove(db.DBObj, index)
			fmt.Println("###################", index, value)
			vehicle.Color = value.Color
			vehicle.Number = value.Number
			vehicle.Slot = value.Slot
			vehicle.Model = value.Model
		}
	}
	db.RW.Unlock()
	if flag {
		fmt.Println("INSIDE DELETE FROM TABLE SLOTS VEHICLES")
		_, err := db.DBmySqlCon.Exec("Delete From slots where car_number = ?", req.Number)
		if err != nil {
			log.Println(err)
			return vehicle, err
		}
		_, err = db.DBmySqlCon.Exec("Delete From vehicles where car_number = ?", req.Number)
		if err != nil {
			log.Println(err)
			return vehicle, err
		}
	}

	if !flag {
		return vehicle, errors.New(constants.NoCarParkedError)
	}
	return vehicle, nil
}

// NewRepository returns instance of Database repository
func NewDatabaseRepository(db *[]Vehicles, dbmySqlCon *sql.DB) repository.DatabaseRepository {

	return &DatabaseImpl{DBObj: db, DBmySqlCon: dbmySqlCon}
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
