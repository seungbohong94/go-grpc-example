package tutorial

import context "context"

// We define a server struct that implements the server interface.
type Server struct {
	UnimplementedTutorialServer
}

func (s *Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{Message: "Hello," + in.GetName()}, nil
}
