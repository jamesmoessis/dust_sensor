package handlers

func RouterHandler(req *Request) (*Response, error) {
	switch req.Path {
	case "/api/settings":
		switch req.Method {
		case "GET":
			return getSettingsHandler(req)
		case "POST":
			return postMeasurementsHandler(req)
		default:
			return &Response{Status: 405}, nil
		}
	case "/api/measurements":
		if req.Method != "POST" {
			return &Response{Status: 405}, nil
		}
		return postMeasurementsHandler(req)
	default:
		return &Response{Status: 404}, nil
	}
}

func updateSettingsHandler(req *Request) (*Response, error) {
	return &Response{}, nil
}

func getSettingsHandler(req *Request) (*Response, error) {
	return &Response{}, nil
}

func postMeasurementsHandler(req *Request) (*Response, error) {
	return &Response{}, nil
}
