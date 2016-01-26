package apihandlers

import "testing"

func TestGetAllSubnets(t *testing.T) {
	subnets := getAllSubnets()
	if subnets == nil {
		t.Error("Error retrieving subnets")
	}
}

func TestGetSubnetIDByTag(t *testing.T) {
	subnetID := getSubnetIDByTag("Name", "Env1as70")
	if subnetID == "" {
		t.Error("Error retrieving subnetID")
	}
}
