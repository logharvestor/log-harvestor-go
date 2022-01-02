package logharvestorgo

type Config struct {
	token    string
	apiUrl   string
	verbose  bool
	batch    bool
	interval int
}

/* NEW */
func New() (*Config, error) {

}
