package logharvestorgo

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

/*
	Forwarder Tests
	* Configuration tests are covered in config_test.go
	===============
	- Initialization
	- Batch Mode Channel
	- Send Log
	- Send Batch Logs
*/

type ForwarderTestSuite struct {
	suite.Suite
	defaultConfig Config
	forwarder     Forwarder
}

func (suite *ForwarderTestSuite) Setup() {
	suite.defaultConfig.token = tokenValid
	suite.defaultConfig.apiUrl = apiUrlValid
	suite.defaultConfig.batch = false
	suite.defaultConfig.verbose = true
	suite.defaultConfig.interval = 30
}

// Initialization
func (suite *ForwarderTestSuite) TestForwarderInit() {
	suite.forwarder = *NewForwarder(suite.defaultConfig)
	f := NewForwarder(suite.defaultConfig)
	// Forwarders should be uniqe
	suite.NotEqual(suite.forwarder, f)
	// Identical Forwarder configs should have equality
	suite.Equal(suite.forwarder.config, f.config)
	// Forwarders should be unique by thier ID
	suite.NotEqual(suite.forwarder.id, f.id)
}

func TestForwarderTestSuite(t *testing.T) {
	suite.Run(t, new(ForwarderTestSuite))
}
