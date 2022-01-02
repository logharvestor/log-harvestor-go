package logharvestorgo

import (
	"testing"
)

/* TEST VARS */
var tokenInvalid = "123ABC"
var tokenValid = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImZvcndhcmRlciJ9.eyJfaWQiOiI2MTI4OTIwYjNjMzQyNTAwMjFkZGQyMTciLCJpYXQiOjE2MzAwNDg3ODN9.sb8lfpp01CC-y0T9Z5XiIEdy-JBeDHSBD8Gd05bZYaQ"
var interval = 30
var testUrl = "http://localhost:3001/api/log"

var defaultConfig = Config{
	token:    "",
	apiUrl:   "",
	verbose:  false,
	batch:    false,
	interval: 10,
}

func TestInit(t *testing.T) {

}
