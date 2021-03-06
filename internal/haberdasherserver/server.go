package haberdasherserver

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/twitchtv/twirp"

	pb "github.com/ezhdanovskiy/twirp-experiments/rpc/haberdasher"
)

// Server implements the Haberdasher service
type Server struct{}

func (s *Server) MakeHat(_ context.Context, size *pb.Size) (hat *pb.Hat, err error) {
	fmt.Printf("haberdasherserver.MakeHat(%+v)\n", size)
	if size.Inches <= 0 {
		return nil, twirp.InvalidArgumentError("inches", "I can't make a hat that small!")
	}

	return &pb.Hat{
		Inches: size.Inches,
		Color:  []string{"white", "black", "brown", "red", "blue"}[rand.Intn(4)],
		Name:   []string{"bowler", "baseball cap", "top hat", "derby"}[rand.Intn(3)],
	}, nil
}
