package logharvestorgo

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var tokenInvalid = "123ABC"
var tokenValid = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImZvcndhcmRlciJ9.eyJfaWQiOiI2MTI4OTIwYjNjMzQyNTAwMjFkZGQyMTciLCJpYXQiOjE2MzAwNDg3ODN9.sb8lfpp01CC-y0T9Z5XiIEdy-JBeDHSBD8Gd05bZYaQ"
var apiUrlInvalid = "tcp://localhost:3001"
var apiUrlValid = "http://localhost:3001"

type ConfigTestSuite struct {
	suite.Suite
	defaultConfig Config
}

type ConfigTableTest struct {
	name     string
	config   Config
	expected bool
}

func (suite *ConfigTestSuite) Setup() {
	suite.defaultConfig.token = Token
	suite.defaultConfig.apiUrl = ApiUrl
	suite.defaultConfig.batch = Batch
	suite.defaultConfig.verbose = Verbose
	suite.defaultConfig.interval = Interval
}

var configTableTests = []ConfigTableTest{
	{"null token & null apiUrl", Config{token: "", apiUrl: "", interval: 0, verbose: false, batch: false}, false},
	{"null token & invalid apiUrl", Config{token: "", apiUrl: apiUrlInvalid, interval: 0, verbose: false, batch: false}, false},
	{"invalid token & null apiUrl", Config{token: tokenInvalid, apiUrl: "", interval: 0, verbose: false, batch: false}, false},
	{"invalid Token & invalid apiUrl", Config{token: tokenInvalid, apiUrl: apiUrlInvalid, interval: 0, verbose: false, batch: false}, false},
	{"valid Token & invalid apiUrl", Config{token: tokenValid, apiUrl: apiUrlInvalid, interval: 20, verbose: false, batch: false}, false},
	{"invalid Token & valid apiUrl", Config{token: tokenInvalid, apiUrl: apiUrlValid, interval: 20, verbose: false, batch: false}, false},
	{"valid Token & valid apiUrl", Config{token: tokenValid, apiUrl: apiUrlValid, interval: 20, verbose: false, batch: false}, true},
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
	suite.True(config.interval == Interval)
	suite.True(config.verbose == Verbose)
	suite.True(config.batch == Batch)
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
