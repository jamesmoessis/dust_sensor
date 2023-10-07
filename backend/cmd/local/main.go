package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jamesmoessis/dust_sensor/backend/handlers"
	"github.com/jamesmoessis/dust_sensor/backend/storage"
)

func main() {
	db := storage.NewDynamoSettingsDb(context.Background())

	h := &localHandler{
		delegate: &handlers.Handler{
			DB: db,
		},
	}

	err := db.CreateSettingsTableIfNotExists(context.Background())
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	http.HandleFunc("/", h.httpHandler)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type localHandler struct {
	delegate *handlers.Handler
}

func (h *localHandler) httpHandler(rw http.ResponseWriter, req *http.Request) {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("err: %v", err)
		rw.WriteHeader(200)
		return
	}
	res, err := h.delegate.RouterHandler(context.Background(), &handlers.Request{
		Body:   string(b),
		Method: req.Method,
		Path:   req.URL.Path,
	})

	if err != nil {
		fmt.Printf("err: %v", err)
		rw.WriteHeader(500)
		return
	}

	rw.WriteHeader(res.Status)
	rw.Write([]byte(res.Body))
}
