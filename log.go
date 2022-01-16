package logharvestorgo

type Log struct {
	Type string      `json:"type" bson:"type"`
	Msg  interface{} `json:"msg" bson:"msg"`
}
type Logs []Log
