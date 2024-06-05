package repositories

import (
	_"errors"
	"fmt"
	"homework/models"
	"sync"
)


type Repository interface {
	GetDevice(string) (models.Device, error)
	CreateDevice(models.Device) error
	DeleteDevice(string) error
	UpdateDevice(models.Device) error
}



type DeviceService struct {
	Repository
}

func NewDeviceService() *DeviceService {
	return &DeviceService{
		Repository: NewRepoDevice(),
	}
}

type RepoDevice struct {
	devices map[string]models.Device
	mu        sync.RWMutex
}

func NewRepoDevice()  *RepoDevice {
	return  &RepoDevice{
		devices: make(map[string]models.Device),
		mu:        sync.RWMutex{},
	}
}


func (ds  *RepoDevice) CreateDevice(device models.Device) error {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	_, ok := ds.devices[device.SerialNum]
	if ok {
		return fmt.Errorf( "%q :%w", device.SerialNum, models.ErrAlredyExist) 
		//&models.ResponseError{Err: errors.New("Device with the same serial number already exist") }
	}
	ds.devices[device.SerialNum] = device
	return nil
}

func (ds  *RepoDevice) GetDevice(serialNumber string) (models.Device, error) {
    ds.mu.Lock()
	defer ds.mu.Unlock()
	device, ok := ds.devices[serialNumber]
	if !ok {
		return models.Device{}, fmt.Errorf( "%q :%w", device.SerialNum, models.ErrNotFound) 
		//&models.ResponseError{Err: errors.New("Device not found") }
	}

	return device, nil
}

func (ds  *RepoDevice) DeleteDevice(serialNumber string) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	_, ok := ds.devices[serialNumber]
	if !ok {
		return fmt.Errorf( "%q :%w", serialNumber, models.ErrNotFound)  
		//&models.ResponseError{Err: errors.New("Device not found") }
	}
	delete(ds.devices, serialNumber)

	return nil
}

func (ds  *RepoDevice) UpdateDevice(device models.Device) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	_, ok := ds.devices[device.SerialNum]
	if !ok {
		return fmt.Errorf( "%q :%w", device.SerialNum, models.ErrNotFound) 
		//&models.ResponseError{Err: errors.New("Device not found") }
	}
	ds.devices[device.SerialNum] = device
	return nil
}