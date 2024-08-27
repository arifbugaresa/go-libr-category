package category

import (
	"context"
	"fmt"
	"go-libr-category/modules/category/pb"
)

type server struct {
	pb.UnimplementedServiceServer
}

func (s *server) InsertCategory(ctx context.Context, req *pb.InsertCategoryRequest) (*pb.InsertCategoryResponse, error) {
	// Implement logic insert category

	fmt.Println(req.Name)
	fmt.Println(req.Description)
	fmt.Println(req.ModifiedBy)

	return &pb.InsertCategoryResponse{Success: true, Message: "Category inserted successfully"}, nil
}
