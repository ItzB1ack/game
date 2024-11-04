package application

import (
	"context"
	"fmt"
	"time"

	"github.com/ItzB1ack/game/pkg/life"
)

type Config struct {
	Width  int
	Height int
}

type Application struct {
	Cfg Config
}

func New(config Config) *Application {
	return &Application{
		Cfg: config,
	}
}

func (a *Application) Run(ctx context.Context) error {
	// Объект для хранения текущего состояния сетки
	currentWorld, err := life.NewWorld(a.Cfg.Height, a.Cfg.Width)
	if err != nil {
		return err
	}
	// Объект для хранения очередного состояния сетки
	nextWorld, err := life.NewWorld(a.Cfg.Height, a.Cfg.Width)
	if err != nil {
		return err
	}
	// Заполняем сетку на 30%
	currentWorld.RandInit(30)
	for {
		// Здесь мы можем записывать текущее состояние  — например, в очередь сообщений. Для нашего примера просто выводим его на экран
		fmt.Println(currentWorld)
		life.NextState(*currentWorld, *nextWorld)
		currentWorld = nextWorld
		// Проверяем контекст
		select {
		case <-ctx.Done():
			return ctx.Err() // Возвращаем причину завершения
		default: // По умолчанию делаем паузу
			time.Sleep(100 * time.Millisecond)
		}
		// Очищаем экран
		fmt.Print("\033[H\033[2J")
	}
}
