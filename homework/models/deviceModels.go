package models

type Device struct {
	SerialNum string `json:"serial_num"`
	Model     string `json:"model"`
	IP        string `json:"ip"`
}
