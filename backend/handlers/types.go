package handlers

import "context"

type Request struct {
	Body   string
	Method string
	Path   string
}

type Response struct {
	Body   string
	Status int
}

type Settings struct {
	IsOn      bool `json:"isOn"`
	Threshold int  `json:"threshold"`
}

type SettingsDB interface {
	GetSettings(context.Context) (Settings, error)
	UpdateSettings(context.Context, Settings) error
}
