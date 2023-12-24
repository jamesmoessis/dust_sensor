package handlers

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type Handler struct {
	DB SettingsDB
}

func (h *Handler) RouterHandler(ctx context.Context, req *Request) (*Response, error) {
	switch req.Path {
	case "/api/settings":
		switch req.Method {
		case "GET":
			return h.getSettingsHandler(ctx, req)
		case "PUT":
			return h.updateSettingsHandler(ctx, req)
		case "OPTIONS":
			return h.allowCorsHandler(ctx, req)
		default:
			return &Response{Status: 405}, nil
		}
	case "/api/measurements":
		if req.Method != "POST" {
			return &Response{Status: 405}, nil
		}
		return h.postMeasurementsHandler(ctx, req)
	default:
		return &Response{Status: 404}, nil
	}
}

func (h *Handler) updateSettingsHandler(ctx context.Context, req *Request) (*Response, error) {
	auth, ok := req.Headers["authorization"]

	if !ok {
		fmt.Println("auth header was not found")
		return &Response{Status: 401, Headers: map[string]string{"WWW-Authenticate": "Basic realm=\"basic\""}}, nil
	}

	base64Str := strings.TrimPrefix(auth, "Basic")
	base64Str = strings.TrimSpace(base64Str)

	plainAuthBytes, err := base64.StdEncoding.DecodeString(base64Str)
	plainAuth := string(plainAuthBytes)
	if err != nil {
		return &Response{
				Status:  400,
				Body:    fmt.Sprintf("invalid base64 in authorization basic header"),
				Headers: map[string]string{"WWW-Authenticate": "Basic realm=\"basic\""}},
			err
	}

	authArr := strings.Split(plainAuth, ":")

	if len(authArr) != 2 {
		return &Response{Status: 400, Body: fmt.Sprintf("invalid auth"), Headers: map[string]string{"WWW-Authenticate": "Basic realm=\"basic\""}}, nil
	}

	userName := authArr[0]
	password := authArr[1]

	hash := sha256.Sum256([]byte(password))
	hashStr := string(hash[:])

	if hashStr != "e4c2a9780a15923de0c007d1b8ee3ee92a6521673a697a23de1c32a782c47938" ||
		userName != "tim" {
		return &Response{
			Status:  401,
			Headers: map[string]string{"WWW-Authenticate": "Basic realm=\"basic\""}}, nil
	}
	settings := &Settings{}

	err = json.Unmarshal([]byte(req.Body), settings)
	if err != nil {
		return &Response{Status: 500}, err
	}

	err = h.DB.UpdateSettings(ctx, *settings)
	if err != nil {
		return &Response{Status: 500}, err
	}

	return &Response{Status: 200, Body: "{\"msg\":\"OK\"}"}, nil
}

func (h *Handler) getSettingsHandler(ctx context.Context, req *Request) (*Response, error) {
	settings, err := h.DB.GetSettings(ctx)
	if err != nil {
		return &Response{Status: 500}, err
	}

	body, err := json.Marshal(settings)
	if err != nil {
		return &Response{Status: 500}, err
	}

	return &Response{
		Status: 200,
		Body:   string(body),
	}, nil
}

func (h *Handler) allowCorsHandler(ctx context.Context, req *Request) (*Response, error) {
	return &Response{
		Status: 200,
	}, nil
}

func (h *Handler) postMeasurementsHandler(ctx context.Context, req *Request) (*Response, error) {
	return &Response{}, nil
}
