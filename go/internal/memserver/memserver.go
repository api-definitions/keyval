// Package memserver provides an implementation of KeyValueService that resides
// in memory.
package memserver

import (
	context "context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/api-definitions/keyval/go/keyvalpb"
)

// Server is a memory-based key/value server.
type Server struct {
	keyvalpb.UnimplementedKeyValueServiceServer
	values map[string][]byte
}

// GetValue implements keyvalpb.KeyValueServiceServer
func (s *Server) GetValue(ctx context.Context, req *keyvalpb.GetValueRequest) (*keyvalpb.Value, error) {
	value, ok := s.values[string(req.GetKey().GetKeyBytes())]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "no entry for key %v", req.GetKey().GetKeyBytes())
	}
	return &keyvalpb.Value{
		ValueBytes: value,
	}, nil
}

// HasValue implements keyvalpb.KeyValueServiceServer
func (s *Server) HasValue(ctx context.Context, req *keyvalpb.HasValueRequest) (*keyvalpb.HasValueResponse, error) {
	_, ok := s.values[string(req.GetKey().GetKeyBytes())]
	return &keyvalpb.HasValueResponse{
		EntryExists: ok,
	}, nil
}

// SetValue implements keyvalpb.KeyValueServiceServer
func (s *Server) SetValue(ctx context.Context, req *keyvalpb.SetValueRequest) (*keyvalpb.SetValueResponse, error) {
	keyStr := string(req.GetKey().GetKeyBytes())
	if req.GetOverwrite() {
		s.values[keyStr] = req.GetValue().GetValueBytes()
	} else {
		_, ok := s.values[string(req.GetKey().GetKeyBytes())]
		if ok {
			return nil, status.Errorf(codes.AlreadyExists, "key/value pair already exists")
		}
		s.values[keyStr] = req.GetValue().GetValueBytes()
	}

	return &keyvalpb.SetValueResponse{}, nil
}

var _ keyvalpb.KeyValueServiceServer = (*Server)(nil)
