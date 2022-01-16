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
	Forwarder     Forwarder
}

// Set Default Configs
func (suite *ForwarderTestSuite) SetupTest() {
	suite.defaultConfig.Token = TokenValid
	suite.defaultConfig.ApiUrl = ApiUrlValid
	suite.defaultConfig.Verbose = false
}

// Init
func (suite *ForwarderTestSuite) TestForwarderInit() {
	suite.Forwarder = *NewForwarder(suite.defaultConfig)
	f := NewForwarder(suite.defaultConfig)
	// Forwarders should be uniqe
	suite.NotEqual(f, suite.Forwarder)
	// Identical Forwarder configs should have equality
	suite.Equal(f.Config, suite.Forwarder.Config)
	// Forwarders should be unique by thier ID
	suite.NotEqual(f.Id, suite.Forwarder.Id)
}

// Init with Verbose mode
func (suite *ForwarderTestSuite) TestVerboseModeInit() {
	// Set Verbose mode to true
	suite.defaultConfig.Verbose = true
	fwdr := *NewForwarder(suite.defaultConfig)
	suite.Forwarder = fwdr
}

// Test Conn - Valid
func (suite *ForwarderTestSuite) TestConnectionValid() {
	fwdr := *NewForwarder(suite.defaultConfig)
	suite.Forwarder = fwdr
	success, msg := suite.Forwarder.testConn()
	suite.Truef(success, msg)
}

// Test Conn - Invalid
func (suite *ForwarderTestSuite) TestConnectionInvalid() {
	suite.defaultConfig.Token = suite.defaultConfig.Token + "asdf"
	fwdr := *NewForwarder(suite.defaultConfig)
	suite.Forwarder = fwdr
	success, msg := suite.Forwarder.testConn()
	suite.Falsef(success, msg)
}

// Send Log
func (suite *ForwarderTestSuite) TestSendLog() {
	suite.Forwarder = *NewForwarder(suite.defaultConfig)
	// Send Test msg
	success, msg := suite.Forwarder.log(Log{Type: "test", Msg: bson.M{"s": 2}})
	// suite.T().Log(msg)
	suite.Truef(success, msg)
}

func TestForwarderTestSuite(t *testing.T) {
	suite.Run(t, new(ForwarderTestSuite))
}
