package logharvestorgo

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var TokenInvalid = "123ABC"
var TokenValid = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImZvcndhcmRlciJ9.eyJfaWQiOiI2MTI4OTIwYjNjMzQyNTAwMjFkZGQyMTciLCJpYXQiOjE2MzAwNDg3ODN9.sb8lfpp01CC-y0T9Z5XiIEdy-JBeDHSBD8Gd05bZYaQ"
var TokenValidLocal = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImZvcndhcmRlciJ9.eyJfaWQiOiI2MDk5Mzg5Mjg4MWQ0MzAwMjkxNzY2MGUiLCJpYXQiOjE2Mjc3MzAzOTZ9.uEY-6s8hK8HX6qy-5Su8Esb-iRXewc9hXYhRLIlALCo"
var ApiUrlInvalid = "tcp://localhost:3001"
var ApiUrlValid = ApiUrl
var ApiUrlValidLocal = "http://localhost:3001/api/log"

type ConfigTestSuite struct {
	suite.Suite
	defaultConfig Config
}

type ConfigTableTest struct {
	name     string
	expected bool
	config   Config
}

func (suite *ConfigTestSuite) Setup() {
	suite.defaultConfig.Token = Token
	suite.defaultConfig.ApiUrl = ApiUrl
	suite.defaultConfig.Verbose = Verbose
}

var configTableTests = []ConfigTableTest{
	{"null Token & null ApiUrl", false, Config{Token: "", ApiUrl: "", Verbose: false}},
	{"null Token & invalid ApiUrl", false, Config{Token: "", ApiUrl: ApiUrlInvalid, Verbose: false}},
	{"invalid Token & null ApiUrl", false, Config{Token: TokenInvalid, ApiUrl: "", Verbose: false}},
	{"invalid Token & invalid ApiUrl", false, Config{Token: TokenInvalid, ApiUrl: ApiUrlInvalid, Verbose: false}},
	{"valid Token & invalid ApiUrl", false, Config{Token: TokenValid, ApiUrl: ApiUrlInvalid, Verbose: false}},
	{"invalid Token & valid ApiUrl", false, Config{Token: TokenInvalid, ApiUrl: ApiUrlValid, Verbose: false}},
	{"valid Token & valid ApiUrl", true, Config{Token: TokenValid, ApiUrl: ApiUrlValid, Verbose: false}},
}

func (suite *ConfigTestSuite) TestConfigsTable() {
	for _, ct := range configTableTests {
		isValid, err := ct.config.Validate()
		suite.T().Logf(`[%v]`, ct.name)
		suite.Equalf(ct.expected, isValid, string(err))
	}
}

func (suite *ConfigTestSuite) TestDefaultConfigFallback() {
	config := NewConfig(Config{})
	suite.True(config.Token == Token)
	suite.True(config.ApiUrl == ApiUrl)
	suite.True(config.Verbose == Verbose)
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
