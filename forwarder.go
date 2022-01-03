package logharvestorgo

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
)

var headers = []string{
	"application/json",
	"application/x-www-form-urlencoded",
}

/* FWDR */
type Forwarder struct {
	id     uuid.UUID
	config Config
	bucket []Log
}

/* INIT */
func NewForwarder(c Config) (*Forwarder, error) {
	return &Forwarder{
		id:     uuid.New(),
		config: c,
		bucket: []Log{},
	}, nil
}

/* FWDR - Init */
func (f *Forwarder) init(c Config) Forwarder {
	return Forwarder{
		config: c,
		bucket: []Log{},
	}
}

/* FWDR - Send Log */
func (f *Forwarder) log(l Log) (bool, string) {
	fmt.Print(l)
	if f.config.batch == true {
		f.bucket = append(f.bucket, l)
	}
	return true, "test"
}

/* FWDR - Test Conn */
func (f *Forwarder) testConn() (bool, string) {
	url := f.config.apiUrl + "/check"
	req, err := http.NewRequest("POST", url, nil)
	req.Header = f.getHeaders()
	if err != nil {
		return false, err.Error()
	}

	client := &http.Client{}
	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return false, err.Error()
	}
	// fmt.Print(string(body))
	return true, string(body)
}

/* UTIL - Build/Get Headers */
func (f *Forwarder) getHeaders() http.Header {
	header := http.Header{}
	header.Add("Content-Type", "application/json")
	header.Add("Content-Type", "application/x-www-form-urlencoded")
	header.Set("Authorization", ("Bearer " + f.config.token))
	return header
}
