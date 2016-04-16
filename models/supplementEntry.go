package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	SupplementEntry struct {
		_id             bson.ObjectId `json:"id"`
		Supplement      string        `json:"supplement"`
		Amount          float64       `json:"amount"`
		MeasurementType string        `json:"measurementType"`
		DateTaken       time.Time     `json:"dateTaken"`
	}
)

type Entries []SupplementEntry
