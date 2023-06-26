package handlers

import "encoding/json"

type Handler struct {
	db SettingsDB
}

func (h *Handler) RouterHandler(req *Request) (*Response, error) {
	switch req.Path {
	case "/api/settings":
		switch req.Method {
		case "GET":
			return h.getSettingsHandler(req)
		case "PUT":
			return h.updateSettingsHandler(req)
		default:
			return &Response{Status: 405}, nil
		}
	case "/api/measurements":
		if req.Method != "POST" {
			return &Response{Status: 405}, nil
		}
		return h.postMeasurementsHandler(req)
	default:
		return &Response{Status: 404}, nil
	}
}

func (h *Handler) updateSettingsHandler(req *Request) (*Response, error) {
	settings := &Settings{}
	err := json.Unmarshal([]byte(req.Body), settings)
	if err != nil {
		return &Response{Status: 500}, err
	}

	err = h.db.UpdateSettings(*settings)
	if err != nil {
		return &Response{Status: 500}, err
	}

	return &Response{Status: 200, Body: "OK"}, nil
}

func (h *Handler) getSettingsHandler(req *Request) (*Response, error) {
	settings, err := h.db.GetSettings()
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

func (h *Handler) postMeasurementsHandler(req *Request) (*Response, error) {
	return &Response{}, nil
}
