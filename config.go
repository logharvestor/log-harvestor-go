package logharvestorgo

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
		token:    Token,
		apiUrl:   ApiUrl,
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
	if c.batch != Batch {
		conf.batch = c.batch
	}
	/* verbose */
	if c.verbose != Verbose {
		conf.verbose = c.verbose || Verbose
	}

	// fmt.Printf("BatchMode: %v\n", conf.batch)
	return conf
}

func (c *Config) validate() (bool, string) {
	/* TOKEN */
	if c.token == "" {
		return false, "Invalid or empty token"
	}

	/* API */
	if c.apiUrl == "" {
		return false, "Invalid or empty token"
	}
	return true, ""
}
