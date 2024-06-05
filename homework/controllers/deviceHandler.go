package controllers

import (
	"encoding/json"
	"homework/models"
	"homework/services"
	"io"
	"net/http"
)

type Handler struct {
	service services.Service
}

func NewHandler(service services.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetDeviceInfo(w http.ResponseWriter, r *http.Request) {
	serialNum := r.URL.Query().Get("serial_num")

	if serialNum == "" {
		w.WriteHeader(http.StatusBadRequest)
		responseBody, _ := json.Marshal(ErrorMessage{Message: "invalid serial number"})
	    _, _ = w.Write(responseBody)
		return
	}

	device, err := h.service.GetDevice(serialNum)
	if err != nil {
		responseBody, _ := json.Marshal(ErrorMessage{Message: err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		_,_ =w.Write(responseBody)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(device)
}

func (h *Handler) CreateDevice(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseBody, _ := json.Marshal(ErrorMessage{Message: "error during reading body"})
		_, _ = w.Write(responseBody)
		return
	}
	var d models.Device
	err = json.Unmarshal(b, &d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseBody, _ := json.Marshal(ErrorMessage{Message: "error during unmarshaling body"})
	    _, _ = w.Write(responseBody)
		return
	}

	respErr := services.ValidateDevice(d)

	if respErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		responseBody, _ := json.Marshal(ErrorMessage{Message: "Invalid date"})
	    _, _ = w.Write(responseBody)
		return
	}

	err = h.service.CreateDevice(d)
	if err != nil {
		responseBody, _ := json.Marshal(ErrorMessage{Message: err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		_,_ = w.Write(responseBody)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) RemoveDevice(w http.ResponseWriter, r *http.Request) {
	serialNum := r.URL.Query().Get("serial_num")

	if serialNum == "" {
		w.WriteHeader(http.StatusBadRequest)
		responseBody, _ := json.Marshal(ErrorMessage{Message: "invalid serial number"})
	    _, _ = w.Write(responseBody)
		return
	}

	err := h.service.DeleteDevice(serialNum)
	if err != nil {
		responseBody, _ := json.Marshal(ErrorMessage{Message: err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		_,_ = w.Write(responseBody)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateDevice(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseBody, _ := json.Marshal(ErrorMessage{Message: "error during reading body"})
	    _, _ = w.Write(responseBody)
		return
	}
	var d models.Device
	err = json.Unmarshal(b, &d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseBody, _ := json.Marshal(ErrorMessage{Message: "error during unmarshaling body"})
	    _, _ = w.Write(responseBody)
		return
	}

	respErr := services.ValidateDevice(d)

	if respErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		responseBody, _ := json.Marshal(ErrorMessage{Message: "Invalid date"})
	    _, _ = w.Write(responseBody)
		return
	}

	err = h.service.UpdateDevice(d)
	if err != nil {
		responseBody, _ := json.Marshal(ErrorMessage{Message: err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		_,_ = w.Write(responseBody)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type ErrorMessage struct {
	Message string `json:"message"`
}