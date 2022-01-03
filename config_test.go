package logharvestorgo

import (
	"fmt"
	"strconv"
	"testing"
)

func FallbackError(constName string, currValue string) string {
	return fmt.Sprintf("Fallback to constant failed for [%+v]: actual value: [%+v]", constName, currValue)
}

func TestDefaultConfigFallback(t *testing.T) {
	config := NewConfig(Config{})
	if config.apiUrl != ApiUrl {
		t.Errorf(FallbackError("ApiUrl", config.apiUrl))
	}
	if config.token != Token {
		t.Errorf(FallbackError("Token", config.token))
	}
	if config.batch != Batch {
		t.Errorf(FallbackError("Batch", strconv.FormatBool(config.batch)))
	}
	if config.interval != Interval {
		t.Errorf(FallbackError("Interval", strconv.Itoa(config.interval)))
	}
	if config.verbose != Verbose {
		t.Errorf(FallbackError("Verbose", strconv.FormatBool(config.verbose)))
	}
}

func TestEmptyConfig(t *testing.T) {
	config := NewConfig(Config{})
	isValid, e := config.validate()
	if !isValid {
		t.Log(e)
	} else {
		t.Errorf("Empty config did not throw error: %+v", config)
	}
}
