package apihandlers

import "testing"

func TestGetAllInstancesInSubnet(t *testing.T) {
	resp := getAllInstancesInSubnet("subnet-9cbcadfe")
	if resp == nil {
		t.Error("Error calling getAllInstancesInSubnet")
	}
}
