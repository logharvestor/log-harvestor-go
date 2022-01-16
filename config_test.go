package logharvestorgo

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var tokenInvalid = "123ABC"
var tokenValid = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImZvcndhcmRlciJ9.eyJfaWQiOiI2MTI4OTIwYjNjMzQyNTAwMjFkZGQyMTciLCJpYXQiOjE2MzAwNDg3ODN9.sb8lfpp01CC-y0T9Z5XiIEdy-JBeDHSBD8Gd05bZYaQ"
var tokenValidLocal = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImZvcndhcmRlciJ9.eyJfaWQiOiI2MDk5Mzg5Mjg4MWQ0MzAwMjkxNzY2MGUiLCJpYXQiOjE2Mjc3MzAzOTZ9.uEY-6s8hK8HX6qy-5Su8Esb-iRXewc9hXYhRLIlALCo"
var apiUrlInvalid = "tcp://localhost:3001"
var apiUrlValid = ApiUrl
var apiUrlValidLocal = "http://localhost:3001/api/log"

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
	suite.defaultConfig.token = Token
	suite.defaultConfig.apiUrl = ApiUrl
	suite.defaultConfig.verbose = Verbose
}

var configTableTests = []ConfigTableTest{
	{"null token & null apiUrl", false, Config{token: "", apiUrl: "", verbose: false}},
	{"null token & invalid apiUrl", false, Config{token: "", apiUrl: apiUrlInvalid, verbose: false}},
	{"invalid token & null apiUrl", false, Config{token: tokenInvalid, apiUrl: "", verbose: false}},
	{"invalid Token & invalid apiUrl", false, Config{token: tokenInvalid, apiUrl: apiUrlInvalid, verbose: false}},
	{"valid Token & invalid apiUrl", false, Config{token: tokenValid, apiUrl: apiUrlInvalid, verbose: false}},
	{"invalid Token & valid apiUrl", false, Config{token: tokenInvalid, apiUrl: apiUrlValid, verbose: false}},
	{"valid Token & valid apiUrl", true, Config{token: tokenValid, apiUrl: apiUrlValid, verbose: false}},
}

func (suite *ConfigTestSuite) TestConfigsTable() {
	for _, ct := range configTableTests {
		isValid, err := ct.config.validate()
		suite.T().Logf(`[%v]`, ct.name)
		suite.Equalf(ct.expected, isValid, string(err))
	}
}

func (suite *ConfigTestSuite) TestDefaultConfigFallback() {
	config := NewConfig(Config{})
	suite.True(config.token == Token)
	suite.True(config.apiUrl == ApiUrl)
	suite.True(config.verbose == Verbose)
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
