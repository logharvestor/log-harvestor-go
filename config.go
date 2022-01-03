package logharvestorgo

import (
	"errors"
)

type Config struct {
	token    string
	apiUrl   string
	verbose  bool
	batch    bool
	interval int
}

/* NEW */
func NewConfig(c Config) *Config {
	conf := &Config{
		apiUrl:   ApiUrl,
		token:    Token,
		interval: Interval,
		verbose:  Verbose,
		batch:    Batch,
	}
	// Fall back to default Constants (./constants.go)
	/* url */
	if c.apiUrl != "" {
		conf.apiUrl = c.apiUrl
	}
	/* token */
	if c.token != "" {
		conf.token = c.token
	}
	/* interval */
	if c.interval != 0 {
		conf.interval = c.interval
	}
	/* batch */
	conf.batch = c.batch || Batch
	/* verbose */
	conf.verbose = c.verbose || Verbose
	return conf
}

func (c *Config) validate() (bool, error) {
	/* TOKEN */
	if c.token == "" {
		return false, errors.New("Invalid or empty token")
	}

	/* API */
	if c.apiUrl == "" {
		return false, errors.New("Invalid or empty token")
	}
	return true, nil
}
