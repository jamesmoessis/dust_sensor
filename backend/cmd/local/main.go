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
	ctx := context.Background()
	db := storage.NewDynamoSettingsDb(ctx)

	exporter, err := handlers.NewExporter(ctx)
	if err != nil {
		log.Fatal(err)
	}
	resource, err := handlers.NewResource("dust_sensor_local", "v1")
	if err != nil {
		log.Fatal(err)
	}
	h := &localHandler{
		delegate: &handlers.Handler{
			DB: db,
			Recorder: handlers.Recorder{
				Resource: *resource,
				Exporter: exporter,
			},
		},
	}

	err = db.CreateSettingsTableIfNotExists(context.Background())
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
	knownValues := []string{
		"threshold", 
		"average", 
		"failurecount", 
		"laptime", 
		"responsetime",
		"maximum",
		"minimum"
	}

	queryParams := make(map[string]string)
	query := req.URL.Query()
	for _, v := range knownValues {
		queryVal := query.Get(v)
		if queryVal != "" {
			queryParams[v] = queryVal
		}
	}
	req.URL.Query()
	res, err := h.delegate.RouterHandler(context.Background(), &handlers.Request{
		Body:        string(b),
		Method:      req.Method,
		Path:        req.URL.Path,
		QueryParams: queryParams,
	})

	if err != nil {
		fmt.Printf("err: %v", err)
		rw.WriteHeader(500)
		return
	}

	reqOrigin := req.Header.Get("Origin")
	resOrigin := "https://dust.jamesmoessis.com"
	if reqOrigin == "http://localhost:3000" {
		resOrigin = reqOrigin
	}

	rw.Header().Add("Access-Control-Allow-Origin", resOrigin)
	rw.Header().Add("Access-Control-Allow-Methods", "GET, PUT")
	rw.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	rw.Header().Add("Content-Type", "application/json")
	if res.Headers != nil {
		for k, v := range res.Headers {
			rw.Header().Add(k, v)
		}
	}

	rw.WriteHeader(res.Status)
	rw.Write([]byte(res.Body))
}
