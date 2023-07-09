package main

import (
	"context"
	"log"

	"github.com/jamesmoessis/dust_sensor/backend/storage"
)

func main() {
	db := storage.NewDynamoSettingsDb(context.Background())
	err := db.CreateSettingsTableIfNotExists(context.Background())
	if err != nil {
		log.Fatalf("err: %v", err)
	}
}
