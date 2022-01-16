package logharvestorgo

import (
	"net/url"
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
	return conf
}

func (c *Config) Validate() (bool, string) {
	/* Token */
	if c.Token == "" {
		return false, "Token not provided"
	}
	if !regexp.MustCompile("^[A-Za-z0-9-_=]+\\.[A-Za-z0-9-_=]+\\.[A-Za-z0-9-_.+/=]*$").MatchString(c.Token) {
		return false, "Token not a valid JWT"
	}

	/* API */
	if c.ApiUrl == "" {
		return false, "Api Url is empty"
	}
	uri, err := url.ParseRequestURI(c.ApiUrl)
	if err != nil {
		return false, "ApiUrl invald"
	}

	switch uri.Scheme {
	case "http":
	case "https":
	default:
		return false, "ApiUrl scheme must be either http or https"
	}

	return true, ""
}
