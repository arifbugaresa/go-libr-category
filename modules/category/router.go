package category

import (
	"go-libr-category/modules/category/pb"
	"google.golang.org/grpc"
)

func Inititator(grpcServer *grpc.Server) {
	pb.RegisterServiceServer(grpcServer, &server{})
}
