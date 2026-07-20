package main

import (
	"testing"
)

func TestHealthJSON(t *testing.T) {
	version := "testing"
	data, err := healthJSON(version, nil)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if string(data) != `{"status":"OK","version":"testing","database_status":"FAIL"}` {
		t.Errorf(`Expected {"status":"OK","version":"testing","database_status":"FAIL"}; got: %v`, string(data))
	}

}
