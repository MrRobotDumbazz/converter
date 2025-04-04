package main

import (
	"context"
	"converter/service"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Convert(value float64, from, to string) float64 {
	return service.Convert(value, from, to)
}

// Метод для получения списка единиц
func (a *App) GetUnits() map[string]float64 {
	return service.AllUnits
}
