package logharvestorgo

type Log struct {
	Typ string      `json:"typ" bson:"typ"`
	Msg  interface{} `json:"msg" bson:"msg"`
}
type Logs []Log
