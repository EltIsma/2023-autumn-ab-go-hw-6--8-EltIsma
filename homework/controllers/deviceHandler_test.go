package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	servMock "homework/controllers/mocks"
	"homework/models"
	"homework/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetDeviceInfo(t *testing.T) {
    mockService := new(servMock.Service)
    ucase := services.NewService(mockService)
    handler := NewHandler(ucase)

    w := httptest.NewRecorder()
    r := httptest.NewRequest(http.MethodGet, "/get?serial_num=123456", nil)
    expectedDevice := models.Device{SerialNum: "123456", Model: "model1", IP: "1.1.1.1"}


    mockService.On("GetDevice", "123456").Return(expectedDevice, nil)

    handler.GetDeviceInfo(w, r)

    assert.Equal(t, http.StatusOK, w.Code)

    var responseDevice models.Device
    err := json.Unmarshal(w.Body.Bytes(), &responseDevice)
    assert.NoError(t, err)
    assert.Equal(t, expectedDevice, responseDevice)


    mockService.AssertExpectations(t)
}

func TestGetDeviceInfo_InvalidSerialNumber(t *testing.T) {
    mockService := new(servMock.Service)
    ucase := services.NewService(mockService)
    handler := NewHandler(ucase)

    w := httptest.NewRecorder()
    r := httptest.NewRequest(http.MethodGet, "/get", nil)

    handler.GetDeviceInfo(w, r)

    assert.Equal(t, http.StatusBadRequest, w.Code)

    var responseBody map[string]string
    err := json.Unmarshal(w.Body.Bytes(), &responseBody)
    assert.NoError(t, err)
    assert.Equal(t, "invalid serial number", responseBody["message"])
}

func TestGetDeviceInfo_DeviceNotFound(t *testing.T) {

    mockService := new(servMock.Service)
    ucase := services.NewService(mockService)
    handler := NewHandler(ucase)

    w := httptest.NewRecorder()
    r := httptest.NewRequest(http.MethodGet, "/get?serial_num=123456", nil)
   
	mockService.On("GetDevice", "123456").Return(models.Device{}, fmt.Errorf("Device not found"))
    handler.GetDeviceInfo(w, r)


    assert.Equal(t, http.StatusBadRequest, w.Code)

    var responseBody map[string]string
    err := json.Unmarshal(w.Body.Bytes(), &responseBody)
    assert.NoError(t, err)
    assert.Equal(t, "Device not found", responseBody["message"])
}

func TestCreateDevice_Success(t *testing.T) {

    mockService := new(servMock.Service)
    ucase := services.NewService(mockService)
    handler := NewHandler(ucase)

    device := models.Device{
		SerialNum: "123456",
		Model:     "model1",
		IP:        "1.1.1.1",
	}
	mockService.On("CreateDevice", device).Return(nil).Times(1)
    

	deviceBytes, _ := json.Marshal(&device)
	rb := bytes.NewReader(deviceBytes)

	w := httptest.NewRecorder()
    r := httptest.NewRequest(http.MethodPost, "/device", rb)

    handler.CreateDevice(w, r)

    assert.Equal(t, http.StatusOK, w.Code)

}

func TestCreateDevice_ErrorDuringReadingBody(t *testing.T) {
   
	mockService := new(servMock.Service)
    ucase := services.NewService(mockService)
    handler := NewHandler(ucase)

	w := httptest.NewRecorder()
    r := httptest.NewRequest(http.MethodPost, "/device", bytes.NewBuffer([]byte(
		[]byte(`
		
          "serial_num": "123456",
		  "model": "model1",
		  "ip": "1.1.1.1"
		}
		`),
	)),)

    handler.CreateDevice(w, r)

    assert.Equal(t, http.StatusInternalServerError, w.Code)
    var responseBody map[string]string
    err := json.Unmarshal(w.Body.Bytes(), &responseBody)
    assert.NoError(t, err)
    assert.Equal(t, "error during unmarshaling body", responseBody["message"])


}

func TestCreateDevice_InvalidDevice(t *testing.T) {

    mockService := new(servMock.Service)
    ucase := services.NewService(mockService)
    handler := NewHandler(ucase)

    w := httptest.NewRecorder()
    r := httptest.NewRequest(http.MethodPost, "/device", bytes.NewBuffer([]byte(
		[]byte(`
		{
          "serial_num": "",
		  "model": "model1",
		  "ip": "1.1.1.1"
		}
		`),
	)),)

    handler.CreateDevice(w, r)

    assert.Equal(t, http.StatusBadRequest, w.Code)

    var responseBody map[string]string
    err := json.Unmarshal(w.Body.Bytes(), &responseBody)
    assert.NoError(t, err)
    assert.Equal(t, "Invalid date", responseBody["message"])
}

func TestCreateDevice_DeviceAlreadyExists(t *testing.T) {

    mockService := new(servMock.Service)
    ucase := services.NewService(mockService)
    handler := NewHandler(ucase)

    w := httptest.NewRecorder()
    r := httptest.NewRequest(http.MethodPost, "/device", bytes.NewBuffer([]byte(
		[]byte(`
		{
          "serial_num": "123456",
		  "model": "model1",
		  "ip": "1.1.1.1"
		}
		`),
	)),)

    mockService.On("CreateDevice", mock.Anything).Return(fmt.Errorf("Device with the same serial number already exist"))

    handler.CreateDevice(w, r)

    assert.Equal(t, http.StatusBadRequest, w.Code)

    var responseBody map[string]string
    err := json.Unmarshal(w.Body.Bytes(), &responseBody)
    assert.NoError(t, err)
    assert.Equal(t, "Device with the same serial number already exist", responseBody["message"])
}

func TestDeleteDevice(t *testing.T) {
    mockService := new(servMock.Service)
    ucase := services.NewService(mockService)
    handler := NewHandler(ucase)

    w := httptest.NewRecorder()
    r := httptest.NewRequest(http.MethodDelete, "/delte?serial_num=123456", nil)

    mockService.On("DeleteDevice","123456").Return(nil)

    handler.RemoveDevice(w, r)

    assert.Equal(t, http.StatusOK, w.Code)

}

func TestDeleteDeviceInvalidSerialNumber(t *testing.T) {

	mockService := new(servMock.Service)
    ucase := services.NewService(mockService)
    handler := NewHandler(ucase)

    w := httptest.NewRecorder()
    r := httptest.NewRequest(http.MethodDelete, "/delete", nil)


    handler.RemoveDevice(w, r)

    assert.Equal(t, http.StatusBadRequest, w.Code)

    var responseBody map[string]string
    err := json.Unmarshal(w.Body.Bytes(), &responseBody)
    assert.NoError(t, err)
    assert.Equal(t, "invalid serial number", responseBody["message"])
}

func TestDeleteDeviceErrorNotFound(t *testing.T) {

    mockService := new(servMock.Service)
    ucase := services.NewService(mockService)
    handler := NewHandler(ucase)


    w := httptest.NewRecorder()
    r := httptest.NewRequest(http.MethodDelete, "/device?serial_num=123456", nil)
	mockService.On("DeleteDevice","123456").Return(fmt.Errorf("Device not found"))

    handler.RemoveDevice(w, r)


    assert.Equal(t, http.StatusBadRequest, w.Code)

    var responseBody map[string]string
    err := json.Unmarshal(w.Body.Bytes(), &responseBody)
    assert.NoError(t, err)
    assert.Equal(t, "Device not found", responseBody["message"])
}

func TestUpdateDevice_Success(t *testing.T) {

    mockService := new(servMock.Service)
    ucase := services.NewService(mockService)
    handler := NewHandler(ucase)

	w := httptest.NewRecorder()
    r := httptest.NewRequest(http.MethodPut, "/device", bytes.NewBuffer([]byte(
		[]byte(`
		{
          "serial_num": "123456",
		  "model": "model1",
		  "ip": "1.1.1.1"
		}
		`),
	)),)


	mockService.On("UpdateDevice", mock.Anything).Return(nil)

    handler.UpdateDevice(w, r)

    assert.Equal(t, http.StatusOK, w.Code)

}

func TestUpdateDeviceErrorBadJson(t *testing.T) {
   
    mockService := new(servMock.Service)
    ucase := services.NewService(mockService)
    handler := NewHandler(ucase)

	w := httptest.NewRecorder()
    r := httptest.NewRequest(http.MethodPut, "/device", bytes.NewBuffer([]byte(
		[]byte(`
		
          "serial_num": "123456",
		  "model": "model1",
		  "ip": "1.1.1.1"
		}
		`),
	)),)

    handler.UpdateDevice(w, r)

    assert.Equal(t, http.StatusInternalServerError, w.Code)
    var responseBody map[string]string
    err := json.Unmarshal(w.Body.Bytes(), &responseBody)
    assert.NoError(t, err)
    assert.Equal(t, "error during unmarshaling body", responseBody["message"])
}

func TestUpdateDevice_InvalidDevice(t *testing.T) {

	mockService := new(servMock.Service)
    ucase := services.NewService(mockService)
    handler := NewHandler(ucase)

    w := httptest.NewRecorder()
    r := httptest.NewRequest(http.MethodPut, "/device", bytes.NewBuffer([]byte(
		[]byte(`
		{
          "serial_num": "",
		  "model": "model1",
		  "ip": "1.1.1.1"
		}
		`),
	)),)

    handler.UpdateDevice(w, r)

    assert.Equal(t, http.StatusBadRequest, w.Code)

    var responseBody map[string]string
    err := json.Unmarshal(w.Body.Bytes(), &responseBody)
    assert.NoError(t, err)
    assert.Equal(t, "Invalid date", responseBody["message"])
}

func TestUpdateDeviceErrorNotFound(t *testing.T) {

    mockService := new(servMock.Service)
    ucase := services.NewService(mockService)
    handler := NewHandler(ucase)

    w := httptest.NewRecorder()
    r := httptest.NewRequest(http.MethodPut, "/device", bytes.NewBuffer([]byte(
		[]byte(`
		{
          "serial_num": "123456",
		  "model": "model1",
		  "ip": "1.1.1.1"
		}
		`),
	)),)


    mockService.On("UpdateDevice", mock.Anything).Return(fmt.Errorf("Device not found"))

    handler.UpdateDevice(w, r)


    assert.Equal(t, http.StatusBadRequest, w.Code)

    var responseBody map[string]string
    err := json.Unmarshal(w.Body.Bytes(), &responseBody)
    assert.NoError(t, err)
    assert.Equal(t, "Device not found", responseBody["message"])
}