package logharvestorgo

import (
	"fmt"

	"github.com/google/uuid"
)

/* FWDR */
type Forwarder struct {
	id     uuid.UUID
	config Config
	bucket []Log
}

/* INIT */
func NewForwarder(c Config) (*Forwarder, error) {

	var _c = Config{}

	return &Forwarder{
		id:     uuid.New(),
		config: _c,
		bucket: []Log{},
	}, nil
}

/* FWDR - Init */
func (f *Forwarder) init(c Config) Forwarder {
	return Forwarder{
		config: c,
		bucket: []Log{},
	}
}

/* FWDR - Send Log */
func (f *Forwarder) log(l Log) (bool, string) {
	fmt.Print(l)
	f.bucket = append(f.bucket, l)
	return true, "test"
}

/* Get Config */
