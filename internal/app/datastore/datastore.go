package datastore

import modbusconnectionmodel "github.com/madllly/go-modbus-desktop-app/internal/app/model/modbus-connection-model"

type DataStore struct {
	Configs modbusconnectionmodel.ConnectionConfig
	Text    string
}

func New() *DataStore {
	return &DataStore{}
}

func (s *DataStore) SetText(text string) {
	s.Text = text
}

func (s *DataStore) GetText() string {
	return s.Text
}
