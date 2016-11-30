package model

import (
	"github.com/ryantomlinson/jaffa-api/db"
)

type Feature struct {
	Id string `db:"id, primarykey, autoincrement" json:"id"`
}

var featuresList = []Feature {
	Feature{Id: "1"},
}

func (f Feature)All() (features []Feature, err error) {
	_, err = db.GetDB().Select(&features, "SELECT id FROM feature")
	return features, err
}