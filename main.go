package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/Rbd3178/task/limiter"
)

func setupControl() FloodControl {
	config := limiter.NewConfig()
	_, err := toml.DecodeFile("configs/floodcontrol.toml", config)
	if err != nil {
		log.Fatal(err)
	}

	return limiter.New(config)
}

func main() {
	fc := setupControl()
	for i := 0; i < 11; i++ {
		ok, err := fc.Check(context.Background(), 1)
		if err != nil {
			fmt.Printf("Check returned an error: %s\n", err)
			continue
		}
		if ok {
			fmt.Printf("Allowed\n")
		} else {
			fmt.Printf("Not allowed\n")
		}
		time.Sleep(1 * time.Second)
	}

}

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}
