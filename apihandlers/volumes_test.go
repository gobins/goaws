package apihandlers

import "testing"

func TestGetAllVolumes(t *testing.T) {
	resp := getAllVolumes()
	if resp == nil {
		t.Error("Error calling getAllVolumes")
	}
}
