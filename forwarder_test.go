package logharvestorgo

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

/*
	Forwarder Tests
	* Configuration tests are covered in config_test.go
	===============
	- Init
	- Init with Batch mode
	- Init with Verbose mode
	- Send Log
	- Send Batch Logs
*/

type ForwarderTestSuite struct {
	suite.Suite
	defaultConfig Config
	forwarder     Forwarder
}

// Set Default Configs
func (suite *ForwarderTestSuite) SetupTest() {
	suite.defaultConfig.token = tokenValid
	suite.defaultConfig.apiUrl = apiUrlValid
	suite.defaultConfig.batch = false
	suite.defaultConfig.verbose = false
	suite.defaultConfig.interval = 1
}

// Init
func (suite *ForwarderTestSuite) TestForwarderInit() {
	suite.forwarder = *NewForwarder(suite.defaultConfig)
	f := NewForwarder(suite.defaultConfig)
	// Forwarders should be uniqe
	suite.NotEqual(f, suite.forwarder)
	// Identical Forwarder configs should have equality
	suite.Equal(f.config, suite.forwarder.config)
	// Forwarders should be unique by thier ID
	suite.NotEqual(f.id, suite.forwarder.id)
}

// Init with Batch mode
func (suite *ForwarderTestSuite) TestBatchModeInit() {
	// Set Batch mode to true
	suite.defaultConfig.batch = true
	fwdr := *NewForwarder(suite.defaultConfig)
	suite.forwarder = fwdr
	// Bucket length should be equal to 0 on Init
	suite.Equal(0, len(suite.forwarder.bucket))
}

// Init with Verbose mode
func (suite *ForwarderTestSuite) TestVerboseModeInit() {
	// Set Verbose mode to true
	suite.defaultConfig.verbose = true
	fwdr := *NewForwarder(suite.defaultConfig)
	suite.forwarder = fwdr
	// TODO
}

// Send Log
func (suite *ForwarderTestSuite) TestSendLog() {

}

// Send Batch Logs
func (suite *ForwarderTestSuite) TestSendBatchLogs() {

}

func TestForwarderTestSuite(t *testing.T) {
	suite.Run(t, new(ForwarderTestSuite))
}
