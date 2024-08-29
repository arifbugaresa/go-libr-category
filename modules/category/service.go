package category

import (
	"context"
	"go-libr-category/modules/category/pb"
)

type service struct {
	pb.UnimplementedServiceServer
	repo Repository
}

func (s *service) InsertCategory(ctx context.Context, req *pb.InsertCategoryRequest) (*pb.InsertCategoryResponse, error) {
	err := s.repo.InsertCategory(ctx, &pb.Category{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	return &pb.InsertCategoryResponse{Success: true, Message: "Category inserted successfully"}, nil
}

func (s *service) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.UpdateCategoryResponse, error) {
	err := s.repo.UpdateCategory(ctx, &pb.Category{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateCategoryResponse{Success: true, Message: "Category updated successfully"}, nil
}

func (s *service) ListCategory(ctx context.Context, req *pb.ListCategoryRequest) (*pb.ListCategoryResponse, error) {
	res, err := s.repo.ListCategory(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.ListCategoryResponse{Categories: res}, nil
}
