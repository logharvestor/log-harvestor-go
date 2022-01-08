package logharvestorgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"

	"github.com/google/uuid"
)

func setInterval(p interface{}, interval time.Duration) chan<- bool {
	ticker := time.NewTicker(interval)
	stopIt := make(chan bool)
	go func() {
		for {

			select {
			case <-stopIt:
				fmt.Println("stop setInterval")
				return
			case <-ticker.C:
				reflect.ValueOf(p).Call([]reflect.Value{})
			}
		}

	}()

	// return the bool channel to use it as a stopper
	return stopIt
}

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

/* START CYCLE */
func (f *Forwarder) startBatchCycle() {
	stopper := setInterval(func() { f.sendBatch(f.bucket) }, 2*time.Second)
	_ = stopper
}

/* INIT */
func NewForwarder(c Config) *Forwarder {
	f := &Forwarder{
		id:     uuid.New(),
		config: c,
		bucket: []Log{},
	}
	// if f.config.batch {
	// 	go func() {
	// 		for now := range time.Tick(time.Second) {
	// 			fmt.Println(now)
	// 			success, msg := f.sendBatch(f.bucket)

	// 			if success {
	// 				print(msg)
	// 			}
	// 		}
	// 	}()
	// }
	return f
}

/* FWDR - Log */
func (f *Forwarder) log(l Log) (bool, string) {
	fmt.Print(l)
	if f.config.batch == true {
		f.bucket = append(f.bucket, l)
	} else {
		f.sendLog(l)
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
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return false, err.Error()
	}
	return true, string(body)
}

/* FWDR - Client Send Batch */
func (f *Forwarder) sendBatch(logs []Log) (bool, string) {
	url := f.config.apiUrl

	data, err := json.Marshal(logs)
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
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return false, err.Error()
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
