package services

import (
	"errors"
	"homework/models"
	repoMock "homework/services/mocks"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDevice(t *testing.T) {
	mockService := new(repoMock.Repository)
	device := models.Device{SerialNum: "123", Model: "model1", IP: "1.1.1.1"}
	mockService.On("CreateDevice", device).Return(nil)

	usecase := NewService(mockService)

	err := usecase.CreateDevice(device)

	assert.NoError(t, err)
}

func TestUpdateDevice(t *testing.T) {
	mockService := new(repoMock.Repository)
	device := models.Device{SerialNum: "123", Model: "model2", IP: "1.1.1.1"}
	mockService.On("UpdateDevice", device).Return(nil)

	usecase := NewService(mockService)

	err := usecase.UpdateDevice(device)

	assert.NoError(t, err)
}

func TestGetDevice(t *testing.T) {
	mockService := new(repoMock.Repository)
	device := models.Device{SerialNum: "123", Model: "model2", IP: "1.1.1.1"}
	mockService.On("GetDevice", device.SerialNum).Return(device, nil)

	usecase := NewService(mockService)

	_, err := usecase.GetDevice(device.SerialNum)

	assert.NoError(t, err)
}

func TestDeleteDevice(t *testing.T) {
	mockService := new(repoMock.Repository)
	device := models.Device{SerialNum: "123", Model: "model1", IP: "1.1.1.1"}
	mockService.On("DeleteDevice", device.SerialNum).Return(nil)

	usecase := NewService(mockService)

	err := usecase.DeleteDevice(device.SerialNum)

	assert.NoError(t, err)
}

// Табличные тесты
func TestValidate(t *testing.T) {

	tests := []struct {
		device      models.Device
		expectedErr error
		//*models.ResponseError
	}{

		{
			models.Device{SerialNum: "12345", Model: "Model1", IP: "192.168.0.1"},
			nil,
		},
		{
			models.Device{SerialNum: "67890", Model: "Model2", IP: "10.0.0.1"},
			nil,
		},
	}
	for _, test := range tests {
		err := ValidateDevice(test.device)
		if err != test.expectedErr {
			t.Errorf("Ошибка при создании устройства с серийным номером %s. Ожидалось: %v, Получено: %v",
				test.device.SerialNum, test.expectedErr, err)
		}
	}
}

func TestValidateError1(t *testing.T) {

	tests := []struct {
		name        string
		device      models.Device
		expectedErr  error
		//*models.ResponseError
	}{

		{
			"invalid ip",
			models.Device{SerialNum: "54321", Model: "Model3", IP: "192.168.0"},
			errors.New("Invalid IP"),
			//&models.ResponseError{Err: errors.New("Invalid IP")},
		},
		{
			"invalid serial number",
			models.Device{SerialNum :"", Model: "Model5", IP: "192.168.0.1"},
			errors.New("Invalid serial number"),
			//&models.ResponseError{Err: errors.New("Invalid serial number")},
		},
		{
			"invalid model",
			models.Device{SerialNum: "54321", Model:"",IP: "10.0.0.1"},
			errors.New("Invalid model"),
			//&models.ResponseError{Err: errors.New("Invalid model")},
		},
	}
	for _, test := range tests {
		err := ValidateDevice(test.device)
		if !reflect.DeepEqual(err, test.expectedErr) {
			t.Errorf("Тест:  %s. Ожидалось: %v, Получено: %v",
				test.name, test.expectedErr, err)
		}
	}
}
