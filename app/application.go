package app

import (
	"access-point/config"
	"access-point/db"
	"access-point/web"
	"access-point/web/utils"
	"sync"
)

type Application struct {
	wg sync.WaitGroup
}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) Init() {
	config.LoadConfig()
	config.GetConfig()
	db.InitDB()
	rabbitmq.InitRabbitMQ()
	utils.InitValidator()
}

func (app *Application) Run() {
	web.StartServer(&app.wg)
}

func (app *Application) Wait() {
	app.wg.Wait()
}

func (app *Application) Cleanup() {
	db.CloseDB()
}