package handlers

type Request struct {
	Body   string
	Method string
	Path   string
}

type Response struct {
	Body   string
	Status int
}

type Handler func(*Request) (*Response, error)
