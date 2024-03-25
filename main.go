package main

import (
	"context"
	"fmt"
	"log"

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
	var id int64
	for {
		fmt.Scanln(&id)
		if id == -1 {
			fmt.Printf("Exiting...\n")
			break
		}
		ok, err := fc.Check(context.Background(), id)
		if err != nil {
			fmt.Printf("Check returned an error: %s\n", err)
		}
		if ok {
			fmt.Printf("Allowed\n")
		} else {
			fmt.Printf("Not allowed\n")
		}
	}

}

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}
