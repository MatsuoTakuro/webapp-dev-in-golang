package ch11

import (
	"fmt"
	"log"
)

func Sub() {
	// list_11_2()
	// list_11_3()
	// list_11_4()
	// list_11_5()
	list_11_6()
}

// implementation detail
type ServiceImple struct{}

func (s *ServiceImple) Apply(id int) error { return nil }

// abstract interface
type OrderService interface {
	Apply(int) error
}

// an user for interface through which he uses the impl.
type Application struct {
	os OrderService
}

func NewApplication(os OrderService) *Application {
	return &Application{
		os: os, // execute DI in initializing an object
	}
}

func (app *Application) Apply(id int) error {
	return app.os.Apply(id)
}

func list_11_2() {
	app := NewApplication(&ServiceImple{})
	err := app.Apply(19)
	if err != nil {
		log.Fatalln("something goes wrong!")
	}
	fmt.Println("successful apply")
}

func (app *Application) SetService(os OrderService) {
	app.os = os // execute DI by call setter method
}

func list_11_3() {
	app := &Application{
		os: nil,
	}
	app.SetService(&ServiceImple{})
	err := app.Apply(19)
	if err != nil {
		log.Fatalln("something goes wrong!")
	}
	fmt.Println("successful apply")
}

func (app *Application) Apply2(os OrderService, id int) error {
	return os.Apply(id) // execute DI in calling method or func with args of implementation detail
}

func list_11_4() {
	app := &Application{
		os: nil,
	}
	err := app.Apply2(&ServiceImple{}, 19)
	if err != nil {
		log.Fatalln("something goes wrong!")
	}
	fmt.Println("successful apply")
}

type EmbeddedApplication struct {
	OrderService
}

func (app *EmbeddedApplication) Run(id int) error {
	return app.Apply(id)
}

func list_11_5() {
	app := &EmbeddedApplication{
		OrderService: &ServiceImple{}, // execute DI in initializing an object
	}
	err := app.Run(19)
	if err != nil {
		log.Fatalln("something goes wrong!")
	}
	fmt.Println("successful apply")
}

func CustomApply(id int) error { return nil }

type ApplicationWithFunc struct {
	Apply func(int) error
}

func (app *ApplicationWithFunc) Run(id int) error {
	return app.Apply(id)
}

func list_11_6() {
	app := &ApplicationWithFunc{
		Apply: CustomApply,
	}
	err := app.Run(19)
	if err != nil {
		log.Fatalln("something goes wrong!")
	}
	fmt.Println("successful apply")
}
