package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockSettingsDb struct {
	shouldError bool
}

var _ SettingsDB = (*MockSettingsDb)(nil)

func (db *MockSettingsDb) GetSettings(_ context.Context) (*Settings, error) {
	var err error
	if db.shouldError {
		err = errors.New("test")
	}
	return &Settings{
		Threshold: 100,
		IsOn:      false,
	}, err
}

func (db *MockSettingsDb) UpdateSettings(_ context.Context, _ Settings) error {
	var err error
	if db.shouldError {
		err = errors.New("test")
	}
	return err
}

func TestGetSettingsHandler(t *testing.T) {
	h := Handler{DB: &MockSettingsDb{}}
	res, err := h.RouterHandler(
		context.Background(),
		&Request{Path: "/api/settings",
			Method: "GET",
		})
	expected := &Response{
		Status: 200,
		Body:   `{"isOn":false,"threshold":100}`,
	}
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestGetSettingsHandlerError(t *testing.T) {
	h := Handler{DB: &MockSettingsDb{shouldError: true}}
	res, err := h.RouterHandler(context.Background(),
		&Request{Path: "/api/settings",
			Method: "GET",
		})
	assert.Error(t, err)
	assert.Equal(t, 500, res.Status)
}

func TestUpdateSettingsHandler(t *testing.T) {
	h := Handler{DB: &MockSettingsDb{}}
	jsonBytes, err := json.Marshal(&Settings{
		Threshold: 2,
		IsOn:      true,
	})
	assert.NoError(t, err)
	jsonString := string(jsonBytes)

	res, err := h.RouterHandler(
		context.Background(),
		&Request{
			Path:   "/api/settings",
			Method: "PUT",
			Body:   jsonString,
		})

	expected := &Response{
		Status: 200,
		Body:   "OK",
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestUpdateSettingsHandlerError(t *testing.T) {
	h := Handler{DB: &MockSettingsDb{shouldError: true}}
	jsonBytes, err := json.Marshal(&Settings{
		Threshold: 2,
		IsOn:      true,
	})
	assert.NoError(t, err)
	jsonString := string(jsonBytes)

	res, err := h.RouterHandler(
		context.Background(),
		&Request{
			Path:   "/api/settings",
			Method: "PUT",
			Body:   jsonString,
		})

	expected := &Response{
		Status: 500,
	}

	assert.Error(t, err)
	assert.Equal(t, expected, res)
}
