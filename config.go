package logharvestorgo

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
)

type Config struct {
	Token   string
	ApiUrl  string
	Verbose bool
}

/* NEW */
func NewConfig(c Config) *Config {
	conf := &Config{
		Token:   Token,
		ApiUrl:  ApiUrl,
		Verbose: Verbose,
	}
	// Fall back to default Constants (./constants.go)
	/* url */
	if c.ApiUrl != "" {
		conf.ApiUrl = c.ApiUrl
	}
	/* Token */
	if c.Token != "" {
		conf.Token = c.Token
	}
	/* Verbose */
	if c.Verbose != Verbose {
		conf.Verbose = c.Verbose || Verbose
	}
	conf.verboseLog(fmt.Sprintf("New Config Created: %+v", conf))
	return conf
}

func (c *Config) Validate() (bool, string) {
	/* Token */
	if c.Token == "" {
		c.verboseLog(fmt.Sprintf("Config.Validate: Failed, Message: %v", "Token not provided"))
		return false, "Token not provided"
	}
	if !regexp.MustCompile("^[A-Za-z0-9-_=]+\\.[A-Za-z0-9-_=]+\\.[A-Za-z0-9-_.+/=]*$").MatchString(c.Token) {
		c.verboseLog(fmt.Sprintf("Config.Validate: Failed, Message: %v", "Token not a valid JWT"))
		return false, "Token not a valid JWT"
	}

	/* API */
	if c.ApiUrl == "" {
		c.verboseLog(fmt.Sprintf("Config.Validate: Failed, Message: %v", "ApiUrl is empty"))
		return false, "ApiUrl is empty"
	}
	uri, err := url.ParseRequestURI(c.ApiUrl)
	if err != nil {
		c.verboseLog(fmt.Sprintf("Config.Validate: Failed, Message: %v", "ApiUrl Invalid"))
		return false, "ApiUrl invald"
	}

	switch uri.Scheme {
	case "http":
	case "https":
	default:
		c.verboseLog(fmt.Sprintf("Config.Validate: Failed, Message: %v", "ApiUrl scheme must be either http or https"))
		return false, "ApiUrl scheme must be either http or https"
	}

	c.verboseLog(fmt.Sprintf("Config.Validate: Success"))
	return true, ""
}

/* UTIL -  VerboseLog */
func (c *Config) verboseLog(msg string) {
	if c.Verbose {
		fmt.Fprintln(os.Stderr, msg)
	}
}
