package logharvestorgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/google/uuid"
)

var headers = []string{
	"application/json",
	"application/x-www-form-urlencoded",
}

/* FWDR */
type Forwarder struct {
	Id            uuid.UUID
	Config        Config
	TotalLogsSent int
}

/* INIT */
func NewForwarder(c Config) *Forwarder {
	f := &Forwarder{
		Id:            uuid.New(),
		Config:        c,
		TotalLogsSent: 0,
	}
	return f
}

// /* FWDR - Log */
func (f *Forwarder) Log(l Log) (bool, string) {
	return f.sendLog(l)
}

/* FWDR - Client Send Log */
func (f *Forwarder) sendLog(l Log) (bool, string) {
	url := f.Config.ApiUrl

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
	f.TotalLogsSent++

	return true, string(body)
}

/* FWDR - Test Conn */
func (f *Forwarder) TestConn() (bool, string) {
	url := f.Config.ApiUrl + "/check"
	req, err := http.NewRequest("POST", url, nil)
	req.Header = f.getHeaders()
	f.verboseLog(fmt.Sprintf("TestConn: %v", url))
	if err != nil {
		f.verboseLog(fmt.Sprintf("TestConn: Failed, ServerResponse: %v", string(req.Response.Status)))
		return false, err.Error()
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		f.verboseLog(fmt.Sprintf("TestConn: Failed, ServerResponse: %v", res.Status))
		return false, err.Error()
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		f.verboseLog(fmt.Sprintf("TestConn: Failed, ServerResponse: %v", res.Status))
		return false, err.Error()
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		f.verboseLog(fmt.Sprintf("TestConn: Failed, ServerResponse: %v", res.Status))
		return false, res.Status
	}
	f.verboseLog(fmt.Sprintf("TestConn: Success, ServerResponse: %v", string(body)))
	return true, string(body)
}

/* UTIL - Build/Get Headers */
func (f *Forwarder) getHeaders() http.Header {
	header := http.Header{}
	header.Add("Content-Type", "application/json")
	header.Add("Content-Type", "application/x-www-form-urlencoded")
	header.Set("Authorization", ("Bearer " + f.Config.Token))
	return header
}

/* UTIL -  VerboseLog */
func (f *Forwarder) verboseLog(msg string) {
	if f.Config.Verbose {
		fmt.Fprintln(os.Stderr, msg)
	}
}
