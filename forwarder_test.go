package logharvestorgo

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
)

/*
	Forwarder Tests
	* Configuration tests are covered in config_test.go
	===============
	- Init
	- Init with Verbose mode
	- Test Conn Valid
	- Test Conn Invalid
	- Send Log
*/

type ForwarderTestSuite struct {
	suite.Suite
	defaultConfig Config
	forwarder     Forwarder
}

// Set Default Configs
func (suite *ForwarderTestSuite) SetupTest() {
	suite.defaultConfig.token = tokenValidLocal
	suite.defaultConfig.apiUrl = apiUrlValid
	suite.defaultConfig.verbose = false
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

// Init with Verbose mode
func (suite *ForwarderTestSuite) TestVerboseModeInit() {
	// Set Verbose mode to true
	suite.defaultConfig.verbose = true
	fwdr := *NewForwarder(suite.defaultConfig)
	suite.forwarder = fwdr
	// TODO
}

// Test Conn - Valid
func (suite *ForwarderTestSuite) TestConnectionValid() {
	fwdr := *NewForwarder(suite.defaultConfig)
	suite.forwarder = fwdr
	success, msg := suite.forwarder.testConn()
	suite.Truef(success, msg)
}

// Test Conn - Invalid
func (suite *ForwarderTestSuite) TestConnectionInvalid() {
	suite.defaultConfig.token = suite.defaultConfig.token + "asdf"
	fwdr := *NewForwarder(suite.defaultConfig)
	suite.forwarder = fwdr
	success, msg := suite.forwarder.testConn()
	suite.Falsef(success, msg)
}

// Send Log
func (suite *ForwarderTestSuite) TestSendLog() {
	// Add prod url
	suite.defaultConfig.token = tokenValidLocal
	suite.forwarder = *NewForwarder(suite.defaultConfig)
	// Send Test msg
	success, msg := suite.forwarder.log(Log{Type: "test", Msg: bson.M{"s": 2}})
	// suite.T().Log(msg)
	suite.Truef(success, msg)
}

func TestForwarderTestSuite(t *testing.T) {
	suite.Run(t, new(ForwarderTestSuite))
}
