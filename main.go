package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/kardianos/service"
	"net/http"
	"os"
)

const serviceName = "SPA web application"
const serviceDescription = "SPA web application"

type program struct{}

func main() {
	fmt.Println(serviceName + " starting...")
	serviceConfig := &service.Config{
		Name:        serviceName,
		DisplayName: serviceName,
		Description: serviceDescription,
	}
	prg := &program{}
	s, err := service.New(prg, serviceConfig)
	if err != nil {
		fmt.Println("Cannot start: " + err.Error())
	}
	err = s.Run()
	if err != nil {
		fmt.Println("Cannot start: " + err.Error())
	}
}

func (p *program) Start(service.Service) error {
	fmt.Println(serviceName + " started")
	go p.run()
	return nil
}

func (p *program) Stop(service.Service) error {
	fmt.Println(serviceName + " stopped")
	return nil
}

func (p *program) run() {
	router := httprouter.New()

	router.ServeFiles("/js/*filepath", http.Dir("js"))
	router.ServeFiles("/css/*filepath", http.Dir("css"))
	router.GET("/", homepage)
	router.POST("/get_content", getContent)
	err := http.ListenAndServe(":80", router)
	if err != nil {
		fmt.Println("Problem starting service: " + err.Error())
		os.Exit(-1)
	}
	fmt.Println(serviceName + " running")
}
