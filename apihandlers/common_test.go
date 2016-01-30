package apihandlers

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
)

func TestUpdateTag(t *testing.T) {
	resp := updateTag("test", "test", []*string{aws.String("subnet-9cbcadfe")})
	if resp == false {
		t.Error("Error updating tag")
	}
}
