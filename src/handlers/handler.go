package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/kalpit-sharma-dev/parkinglot-service/src/constants"
	"github.com/kalpit-sharma-dev/parkinglot-service/src/service"

	errs "github.com/kalpit-sharma-dev/parkinglot-service/src/errors"
	"github.com/kalpit-sharma-dev/parkinglot-service/src/models"
)

type Handler struct {
	service service.ParkingLotService
}

const (
	carColor  = "color"
	carNumber = "number"
)

func (h *Handler) HandleCreateSlotEvent(w http.ResponseWriter, r *http.Request) {
	log.Println("INFO : HandleCreateParkEvent", r)
	decoder := json.NewDecoder(r.Body)
	var req models.Slot
	err := decoder.Decode(&req)
	if err != nil {
		log.Println("unable to decode", err)
		w.WriteHeader(http.StatusBadRequest)
		errs.HandleError(w, err)
		return
	}
	slotEvent, err := h.service.CreateSlotEvent(r.Context(), req)
	if err != nil {
		log.Println("error", err)
		w.WriteHeader(http.StatusInternalServerError)
		errs.HandleError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(slotEvent)
}

func (h *Handler) HandleGetGovernmentAPI(w http.ResponseWriter, r *http.Request) {
	paramColor := r.URL.Query().Get(carColor)
	paramNumber := r.URL.Query().Get(carNumber)
	if len(strings.TrimSpace(paramColor)) != 0 {
		h.HandleGetAllCarsWithColor(w, r)
		return
	} else if len(strings.TrimSpace(paramNumber)) != 0 {
		h.HandleGetSlotNumberWithCarID(w, r)
		return
	}

}

func (h *Handler) HandleGetAllCarsWithColor(w http.ResponseWriter, r *http.Request) {
	paramColor := r.URL.Query().Get(carColor)

	resp, err := h.service.GetAllCarsWithColor(r.Context(), paramColor)
	if err != nil {
		log.Println("error", err)
		errs.HandleError(w, err)
		return
	}
	json.NewEncoder(w).Encode(resp)

}

func (h *Handler) HandleGetAllSlotNumberWithColor(w http.ResponseWriter, r *http.Request) {
	paramColor := r.URL.Query().Get(carColor)
	resp, err := h.service.GetAllSlotNumberWithColor(r.Context(), paramColor)
	if err != nil {
		log.Println("error", err)
		errs.HandleError(w, err)
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

	paramNumber := r.URL.Query().Get(carNumber)
	resp, err := h.service.GetSlotNumberWithCarID(r.Context(), paramNumber)
	if err != nil {
		customError := errs.ParkingCustomError{
			ID:      400,
			Message: constants.NoCarParkedError,
		}
		log.Println("error", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customError)
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
		customError := errs.ParkingCustomError{
			ID:      500,
			Message: constants.NoSlotError,
		}
		log.Println("error", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(customError)
		//errs.HandleError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(parkEvent)
}

func (h *Handler) HandleExitParkEvent(w http.ResponseWriter, r *http.Request) {

	var req models.Vehicle

	log.Println("INFO : HandleExitParkEvent", r)

	paramNumber := r.URL.Query().Get(carNumber)

	req.Number = paramNumber
	event, err := h.service.ExitParkEvent(r.Context(), req)
	if err != nil {
		customError := errs.ParkingCustomError{
			ID:      400,
			Message: constants.NoCarParkedError,
		}
		log.Println("error in exit park event", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}
func NewHandler(service service.ParkingLotService) *Handler {
	return &Handler{
		service: service,
	}
}
