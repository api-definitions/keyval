package examplelib

import (
	"github.com/golang/glog"

	"github.com/api-definitions/keyval/proto/go/keyvalpb"
)

var X = 5

func ComputeSomthing() int {
	req := keyvalpb.GetValueRequest{
		Key: &keyvalpb.Key{},
	}
	glog.Infof("")
	return 3
}
