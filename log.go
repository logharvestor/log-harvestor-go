package logharvestorgo

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Log struct {
	Type string `json:"type" bson:"type"`
	Msg  bson.M `json:"msg" bson:"msg"`
}
type Logs []Log
