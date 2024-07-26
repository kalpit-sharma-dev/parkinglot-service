package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/kalpit-sharma-dev/parkinglot-service/src/handlers"
	"github.com/kalpit-sharma-dev/parkinglot-service/src/repository/db"
	"github.com/kalpit-sharma-dev/parkinglot-service/src/service"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func LoadRoute() {
	log.Println("INFO : Loading Router")
	router := mux.NewRouter().PathPrefix("/parkinglot-service/api").Subrouter()
	router.Use(headerMiddleware)
	registerAppRoutes(router)
	log.Println("INFO : Router Loaded Successfully")
	log.Println("INFO : Application is started Successfully")
	// wg := &sync.WaitGroup{}
	// go LoadSlots(wg)
	http.ListenAndServe(":9999", router)
}

var SlotCounter uint64

func registerAppRoutes(r *mux.Router) {
	log.Println("INFO : Registering Router ")

	log.Println("INFO : Registering Router ")

	var err error
	// Connect to MySQL database
	dbmySqlCon, err := sql.Open("mysql", "kalpit:password@tcp(192.168.100.4:3306)/demo")
	if err != nil {
		log.Fatal(err)
	}

	// Test database connection
	if err := dbmySqlCon.Ping(); err != nil {
		log.Fatal(err)
	}

	var dbConn db.DatabaseImpl

	vehicleDB := make([]db.Vehicles, 1)
	dbConn.DBObj = &vehicleDB

	eventRepo := db.NewDatabaseRepository(dbConn.DBObj, dbmySqlCon)

	eventService := service.NewParkingLotService(eventRepo)

	eventHandlers := handlers.NewHandler(eventService)

	r.HandleFunc("/parking/slots", eventHandlers.HandleCreateSlotEvent).Methods(http.MethodPost)

	r.HandleFunc("/parking/cars", eventHandlers.HandleGetGovernmentAPI).Methods(http.MethodGet) //quer param color,number

	r.HandleFunc("/parking/slots", eventHandlers.HandleGetAllSlotNumberWithColor).Methods(http.MethodGet) ////quer param color return slot numbers

	r.HandleFunc("/parking/vehicle/park", eventHandlers.HandleCreateParkEvent).Methods(http.MethodPost)
	r.HandleFunc("/parking/vehicle/exit", eventHandlers.HandleExitParkEvent).Methods(http.MethodGet)

	// /r.HandleFunc("/parking/slots/status", eventHandlers.HandleCreateEvents).Methods(http.MethodPost)

	log.Println("INFO : Router Registered Successfully")
}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
