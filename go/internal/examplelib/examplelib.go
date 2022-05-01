package examplelib

import (
	"github.com/golang/glog"

	"github.com/api-definitions/keyval/proto/go/keyvalpb"
	//"github.com/api-definitions/keyval/bazel-bin/proto/keyvalpb_go_proto_/github.com/api-definitions/keyval/proto/go/keyvalpb"
	//"github.com/api-definitions/keyval/bazel-bin/proto/keyvalpb_go_proto_/github.com/api-definitions/keyval/go/proto/keyvalpb"
	//"github.com/api-definitions/keyval/bazel-bin/proto/keyvalpb_go_proto_/github.com/api-definitions/keyval/proto/go/keyvalpb"
)

var X = 5

func ComputeSomthing() int {
	keyvalpb.GetValueRequest{
		Key: &keyvalpb.Key{},
	}
	req := keyvalpb.GetValueRequest{
		Key: &keyvalpb.Key{},
	}
	glog.Infof("%v", req)
	return 3
}
