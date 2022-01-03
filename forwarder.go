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
	return &Forwarder{
		id:     uuid.New(),
		config: c,
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
	if f.config.batch == true {
		f.bucket = append(f.bucket, l)
	}
	return true, "test"
}
