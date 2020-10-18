package butlerserver

import (
	"context"

	pb "github.com/ezhdanovskiy/twirp-experiments/rpc/butler"
)

// Server implements the Butler service
type Server struct{}

func (s *Server) Hello(_ context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{Text: "Hello " + req.Subject}, nil
}
