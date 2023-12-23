package handlers

import (
	"context"
	"encoding/json"
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

	// Pick a username / password
	// take echo -n "mypassword" | shasum -a 256
	// make the code take the password's hash (use the golang sha256 lib)
	// There should be a header incoming that looks like:
	// Authorization: Basic mypassword
	// Then take the sha 256 password hash and compare it to what is in the database.

	// If it doesn't have the correct password, then it needs to return
	// a response &Response {Status: 401, Headers: {"WWW-Authenticate": "Basic realm=\"basic\""}}

	// Store the hash in a variable at first (and we will put it in the database later, maybe change it)

	settings := &Settings{}
	err := json.Unmarshal([]byte(req.Body), settings)
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
