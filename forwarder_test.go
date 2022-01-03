package logharvestorgo

import (
	"testing"
)

/* TEST VARS */
var tokenInvalid = "123ABC"

// var tokenValid = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImZvcndhcmRlciJ9.eyJfaWQiOiI2MTI4OTIwYjNjMzQyNTAwMjFkZGQyMTciLCJpYXQiOjE2MzAwNDg3ODN9.sb8lfpp01CC-y0T9Z5XiIEdy-JBeDHSBD8Gd05bZYaQ"
var tokenValid = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImZvcndhcmRlciJ9.eyJfaWQiOiI2MDk5Mzg5Mjg4MWQ0MzAwMjkxNzY2MGUiLCJpYXQiOjE2Mjc3MzAzOTZ9.uEY-6s8hK8HX6qy-5Su8Esb-iRXewc9hXYhRLIlALCo"
var interval = 30
var testUrl = "http://localhost:3001/api/log"

var defaultConfig = NewConfig(Config{})

func TestDefaultInit(t *testing.T) {
	forwarder, e := NewForwarder(*defaultConfig)
	if e != nil {
		t.Error(e, forwarder)
	}
}

func TestSendLogNoBatch(t *testing.T) {
	conf := NewConfig(Config{token: tokenValid})
	forwarder, e := NewForwarder(*conf)
	if e != nil {
		t.Error(e, forwarder)
	}
	success, msg := forwarder.log(Log{Lvl: "test", Msg: "{s: 1}"})
	if !success {
		t.Error(msg)
	}
	if len(forwarder.bucket) != 0 {
		t.Errorf("Log appended to bucket while not running in BATCH - Bucket: %+v", forwarder.bucket)
	}
}

func TestSendLogBatch(t *testing.T) {
	conf := NewConfig(Config{token: tokenValid, batch: true})
	forwarder, e := NewForwarder(*conf)
	if e != nil {
		t.Error(e, forwarder)
	}
	success, msg := forwarder.log(Log{Lvl: "test", Msg: "{s: 1}"})
	if !success {
		t.Error(msg)
	}
	if len(forwarder.bucket) == 0 {
		t.Errorf("Log failed to append to bucket while running in BATCH - Bucket: %+v", forwarder.bucket)
	}
}

func TestConn(t *testing.T) {
	conf := NewConfig(Config{token: tokenValid, apiUrl: testUrl})
	forwarder, e := NewForwarder(*conf)
	if e != nil {
		t.Error(e, forwarder)
	}
	success, msg := forwarder.testConn()

	if !success {
		t.Error(msg)
	} else {
		t.Log(msg)
	}
}
