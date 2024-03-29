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

type forwarderTableTest struct {
	name     string
	expected bool
	log      Log
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
	success, msg := suite.Forwarder.TestConn()
	suite.Truef(success, msg)
}

// Test Conn - Invalid
func (suite *ForwarderTestSuite) TestConnectionInvalid() {
	suite.defaultConfig.Token = suite.defaultConfig.Token + "asdf"
	fwdr := *NewForwarder(suite.defaultConfig)
	suite.Forwarder = fwdr
	success, msg := suite.Forwarder.TestConn()
	suite.Falsef(success, msg)
}

// Send Log
func (suite *ForwarderTestSuite) TestSendLog() {
	suite.Forwarder = *NewForwarder(suite.defaultConfig)
	success, msg := suite.Forwarder.Log(Log{Typ: "test", Msg: bson.M{"s": 2}})
	suite.Truef(success, msg)
}

var forwarderTableTests = []forwarderTableTest{
	{"Single String", true, Log{Typ: "test", Msg: "hello world!"}},
	{"Single Number", true, Log{Typ: "test", Msg: 123}},
	{"Single Object", true, Log{Typ: "test", Msg: bson.M{"a": 123, "b": "456", "c": "hello"}}},
	{"Nested Object", true, Log{Typ: "test", Msg: bson.M{"a": bson.M{"b": "123"}}}},
	{"Nested Mixed", true, Log{Typ: "test", Msg: bson.M{"a": 123, "b": bson.M{"c": "123", "d": nil}}}},
	{"String Array", true, Log{Typ: "test", Msg: bson.A{"hello", "mars", "goodbye", "world"}}},
	{"Number Array", true, Log{Typ: "test", Msg: bson.A{1, 2, 34, 567, 8, 90}}},
	{"Object Array", true, Log{Typ: "test", Msg: bson.A{bson.M{"question": "Hello?", "answer": "World!"}, bson.M{
		"question": "So long?", "answer": "Thanks for all the fish!"}}}},
	{"Mixed Array", true, Log{Typ: "test", Msg: bson.A{123, "abc", bson.A{1, 2, 3}, bson.M{"question": "Hello?", "answer": "World!"}, bson.M{
		"question": "So long?", "answer": "Thanks for all the fish!"}}}},
	{"Empty Msg", true, Log{Typ: "test", Msg: ""}},
	{"Nil Msg", false, Log{Typ: "test", Msg: nil}},
}

func (suite *ForwarderTestSuite) TestForwardersTable() {
	for _, ft := range forwarderTableTests {
		success, msg := suite.Forwarder.Log(ft.log)
		suite.T().Logf(`[%v]`, ft.name)
		suite.Equalf(ft.expected, success, string(msg))
	}
}

func TestForwarderTestSuite(t *testing.T) {
	suite.Run(t, new(ForwarderTestSuite))
}
