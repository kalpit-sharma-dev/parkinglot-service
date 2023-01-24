package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kalpit-sharma-dev/parkinglot-service/src/service"

	"github.com/kalpit-sharma-dev/parkinglot-service/src/errors"
	"github.com/kalpit-sharma-dev/parkinglot-service/src/models"
)

type Handler struct {
	service service.ParkingLotService
}

const (
	carColor  = "color"
	carNumber = "number"
)

func (h *Handler) HandleGetAllCarsWithColor(w http.ResponseWriter, r *http.Request) {
	paramColor := r.URL.Query().Get(carColor)

	resp, err := h.service.GetAllCarsWithColor(r.Context(), paramColor)
	if err != nil {
		log.Println("error", err)
		errors.HandleError(w, err)
		return
	}
	json.NewEncoder(w).Encode(resp)

}

func (h *Handler) HandleGetAllSlotNumberWithColor(w http.ResponseWriter, r *http.Request) {
	paramColor := r.URL.Query().Get(carColor)
	resp, err := h.service.GetAllSlotNumberWithColor(r.Context(), paramColor)
	if err != nil {
		log.Println("error", err)
		errors.HandleError(w, err)
		return
	}

	fmt.Println(resp)
	if len(resp) < 1 {
		log.Println(err, "No Slot found for this color : ", paramColor)
		w.WriteHeader(http.StatusOK)

		resp = make([]models.Vehicle, 0)
		json.NewEncoder(w).Encode(resp)
		return
	}
	json.NewEncoder(w).Encode(resp)

}
func (h *Handler) HandleGetSlotNumberWithCarID(w http.ResponseWriter, r *http.Request) {

	//params := mux.Vars(r)
	//paramNumber := params[carNumber]

	paramNumber := r.URL.Query().Get(carNumber)
	resp, err := h.service.GetSlotNumberWithCarID(r.Context(), paramNumber)
	if err != nil {
		log.Println("error", err)
		errors.HandleError(w, err)
		return
	}
	json.NewEncoder(w).Encode(resp)

}
func (h *Handler) HandleCreateParkEvent(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var req models.Vehicle
	err := decoder.Decode(&req)
	if err != nil {
		log.Println("unable to decode", err)
		return
	}
	log.Println("INFO : HandleCreateParkEvent", r)

	parkEvent, err := h.service.CreateParkEvent(r.Context(), req)
	if err != nil {
		log.Println("error", err)
		w.WriteHeader(http.StatusInternalServerError)
		errors.HandleError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(parkEvent)
}

func (h *Handler) HandleExitParkEvent(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var req models.Vehicle
	err := decoder.Decode(&req)
	if err != nil {
		log.Println("unable to decode", err)
		return
	}
	log.Println("INFO : HandleExitParkEvent", r)
	// pathParam := mux.Vars(r)
	// number := pathParam[carNumber]

	paramNumber := r.URL.Query().Get(carNumber)

	req.Number = paramNumber
	event, err := h.service.ExitParkEvent(r.Context(), req)
	if err != nil {
		log.Println("error", err)
		w.WriteHeader(http.StatusInternalServerError)
		errors.HandleError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}
func NewHandler(service service.ParkingLotService) *Handler {
	return &Handler{
		service: service,
	}
}
