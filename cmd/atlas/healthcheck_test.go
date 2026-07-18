package main

import (
	"testing"
)

func TestHealthJSON(t *testing.T) {
	version := "1.2.3"
	data, err := healthJSON(version)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if string(data) != `{"status":"OK","version":"1.2.3"}` {
		t.Errorf("Expected Status: OK and Version: 1.2.3; got: %v", string(data))
	}

}
