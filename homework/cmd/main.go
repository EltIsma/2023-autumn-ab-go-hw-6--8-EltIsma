package main

import (
	"fmt"
	"homework/controllers"
	"homework/repositories"
	"homework/services"
	"log"
	"net/http"
	"os"
)

func main() {

	addr, check := os.LookupEnv("ADDRESS")
	if !check{
		addr = "127.0.0.1"
	}
	port, check:= os.LookupEnv("PORT")
	if !check{
		port = "8080"
	}
	repo := repositories.NewDeviceService()
	service := services.NewService(repo)
	handler := controllers.NewHandler(service)
	http.HandleFunc("/get", handler.GetDeviceInfo)
	http.HandleFunc("/create", handler.CreateDevice)
	http.HandleFunc("/update", handler.UpdateDevice)
	http.HandleFunc("/delete", handler.RemoveDevice)
	log.Printf("Starting server on %s:%s", addr, port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", addr, port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
