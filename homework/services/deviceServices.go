package services

import (
	"errors"
	"homework/models"
	"net"
)

type Service interface {
	GetDevice(string) (models.Device, error)
	CreateDevice(models.Device) error
	DeleteDevice(string) error
	UpdateDevice(models.Device) error
}

type Usercase struct {
	devices Service
}


func NewService(devices Service) *Usercase {
	return &Usercase{
		devices: devices,
	}
}


func (u *Usercase) CreateDevice(device models.Device) (error) {
	return u.devices.CreateDevice(device)
}

func (u *Usercase) GetDevice(serialNumber string) (models.Device, error) {
	return u.devices.GetDevice(serialNumber)
}

func (u *Usercase) DeleteDevice(serialNumber string) (error) {
	return u.devices.DeleteDevice(serialNumber)
}

func (u *Usercase) UpdateDevice(device models.Device) (error) {
	return u.devices.UpdateDevice(device)
}

func ValidateDevice(d models.Device) error {
	if d.SerialNum == "" {
		return errors.New("Invalid serial number")
		//&models.ResponseError{Err: errors.New("Invalid serial number") }
	}

	if d.Model == "" {
		return errors.New("Invalid model")
		//&models.ResponseError{Err: errors.New("Invalid model") }
	}

	if d.IP == "" || !checkIPv4(d.IP) {
		return errors.New("Invalid IP")
		//&models.ResponseError{Err: errors.New("Invalid IP") }
	}

	return nil
}



func checkIPv4(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil && parsedIP.To4() != nil
}
