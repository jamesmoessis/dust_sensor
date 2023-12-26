package handlers

import "context"

type Request struct {
	Body        string
	Method      string
	Path        string
	Headers     map[string]string
	QueryParams map[string]string
}

type Response struct {
	Body    string
	Status  int
	Headers map[string]string
}

type Settings struct {
	IsOn      bool `json:"isOn"`
	Threshold int  `json:"threshold"`
}

type SettingsDB interface {
	GetSettings(context.Context) (*Settings, error)
	UpdateSettings(context.Context, Settings) error
}
