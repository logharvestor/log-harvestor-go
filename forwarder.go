package logharvestorgo

import (
	"bytes"
	"encoding/json"
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
	id            uuid.UUID
	config        Config
	totalLogsSent int
}

/* INIT */
func NewForwarder(c Config) *Forwarder {
	f := &Forwarder{
		id:            uuid.New(),
		config:        c,
		totalLogsSent: 0,
	}
	return f
}

// /* FWDR - Log */
func (f *Forwarder) log(l Log) (bool, string) {
	return f.sendLog(l)
}

/* FWDR - Client Send Log */
func (f *Forwarder) sendLog(l Log) (bool, string) {
	url := f.config.apiUrl

	data, err := json.Marshal(l)
	if err != nil {
		return false, err.Error()
	}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(string(data)))
	req.Header = f.getHeaders()
	if err != nil {
		return false, err.Error()
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return false, err.Error()
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err.Error()
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return false, res.Status
	}
	f.totalLogsSent++

	return true, string(body)
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
	if err != nil {
		return false, err.Error()
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err.Error()
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return false, res.Status
	}
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
