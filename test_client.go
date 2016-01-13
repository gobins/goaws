package client

import (
	"testing"
)

func TestGetEC2Session(t *testing.T) {
	session, err := client.getSession()
	if err != nil {
		t.Error("Error getting a session")
	}
}
