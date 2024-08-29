package category

import (
	"database/sql"
	"go-libr-category/modules/category/pb"
	"google.golang.org/grpc"
)

func Inititator(grpcServer *grpc.Server, connection *sql.DB) {
	repo := NewRepository(connection)
	pb.RegisterServiceServer(grpcServer, &service{repo: repo})
}
