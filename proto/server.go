package proto

import "context"

type Server struct {

}

//func (s *Server) mustEmbedUnimplementedGreeterServer() {
//	panic("implement me")
//}

func (s *Server) SayHello(_ context.Context,in *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: "Hello From the Server! -- " + in.Name}, nil
}
