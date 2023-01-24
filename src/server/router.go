package server

import (
	"log"
	"net/http"

	"github.com/kalpit-sharma-dev/parkinglot-service/src/handlers"
	"github.com/kalpit-sharma-dev/parkinglot-service/src/repository/db"
	"github.com/kalpit-sharma-dev/parkinglot-service/src/service"

	"github.com/gorilla/mux"
)

func LoadRoute() {
	log.Println("INFO : Loading Router")
	router := mux.NewRouter().PathPrefix("/parkinglot-service/api").Subrouter()
	router.Use(headerMiddleware)
	registerAppRoutes(router)
	log.Println("INFO : Router Loaded Successfully")
	log.Println("INFO : Application is started Successfully")
	http.ListenAndServe(":9999", router)
}

func registerAppRoutes(r *mux.Router) {
	log.Println("INFO : Registering Router ")

	log.Println("INFO : Registering Router ")

	var dbConn db.DatabaseImpl
	// /c, err := db.GetDatabaseProvider(DBConfig.DBUser, DBConfig.DBPassword, DBConfig.DBName)
	vehicleDB := make([]db.Vehicles, 100000)
	dbConn.DBObj = &vehicleDB

	eventRepo := db.NewDatabaseRepository(dbConn.DBObj)

	eventService := service.NewParkingLotService(eventRepo)

	eventHandlers := handlers.NewHandler(eventService)

	r.HandleFunc("/parking/slots", eventHandlers.HandleGetAllCarsWithColor).Methods(http.MethodPost)

	r.HandleFunc("/parking/cars", eventHandlers.HandleGetAllCarsWithColor).Methods(http.MethodPut)         //quer param color
	r.HandleFunc("/parking/cars/{id}", eventHandlers.HandleGetSlotNumberWithCarID).Methods(http.MethodGet) //return slot number
	r.HandleFunc("/parking/slots", eventHandlers.HandleGetAllSlotNumberWithColor).Methods(http.MethodGet)  ////quer param color return slot numbers

	r.HandleFunc("/parking/vehicle/park", eventHandlers.HandleCreateParkEvent).Methods(http.MethodPost)
	r.HandleFunc("/parking/vehicle/exit", eventHandlers.HandleExitParkEvent).Methods(http.MethodPost)

	// /r.HandleFunc("/parking/slots/status", eventHandlers.HandleCreateEvents).Methods(http.MethodPost)

	log.Println("INFO : Router Registered Successfully")
}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
