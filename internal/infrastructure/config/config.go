package config

import (
	"github.com/bulutcan99/weekly-task-scheduler/internal/infrastructure/env"
)

var (
	Host       = &env.Env.Host
	ServerPort = &env.Env.ServerPort
	DbPort     = &env.Env.DbPort
	DbName     = &env.Env.DbName
)

type (
	Container struct {
		App   *App
		Mongo *Mongo
		Fiber *Fiber
	}

	App struct {
		Name string
	}

	Mongo struct {
		Host string
		Port int
		Name string
	}

	Fiber struct {
		Host string
		Port int
	}
)

func New() *Container {
	app := &App{
		Name: "scheduler",
	}

	mongo := &Mongo{
		Host: *Host,
		Port: *DbPort,
		Name: *DbName,
	}

	fiber := &Fiber{
		Host: *Host,
		Port: *ServerPort,
	}

	return &Container{
		App:   app,
		Mongo: mongo,
		Fiber: fiber,
	}
}
