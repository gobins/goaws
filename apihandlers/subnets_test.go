package apihandlers

import "testing"

func TestGetSubnets(t *testing.T) {
	subnets := getSubnets("ap-southeast-2")
	if subnets == nil {
		t.Error("Error retrieving subnets")
	}
}

func TestParseSubnetsData(t *testing.T) {

}
