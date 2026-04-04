package models

import "time"

type Metric struct {
	Key       string    `json:"key"`
	Label     string    `json:"label"`
	Value     float64   `json:"value"`
	Unit      string    `json:"unit"`
	Period    string    `json:"period"`
	UpdatedAt time.Time `json:"updated_at"`
}
