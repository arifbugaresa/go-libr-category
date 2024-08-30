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
	err := s.repo.InsertCategory(ctx, DTOCategory{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	return &pb.InsertCategoryResponse{Success: true, Message: "Category inserted successfully"}, nil
}

func (s *service) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.UpdateCategoryResponse, error) {
	err := s.repo.UpdateCategory(ctx, DTOCategory{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateCategoryResponse{Success: true, Message: "Category updated successfully"}, nil
}

func (s *service) ListCategory(ctx context.Context, req *pb.ListCategoryRequest) (*pb.ListCategoryResponse, error) {
	categories, err := s.repo.ListCategory(ctx)
	if err != nil {
		return nil, err
	}

	// convert into response pb
	categoryResponse := make([]*pb.Category, len(categories))
	for _, category := range categories {
		categoryResponse = append(categoryResponse, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return &pb.ListCategoryResponse{Categories: categoryResponse}, nil
}

func (s *service) GetCategoryById(ctx context.Context, req *pb.GetCategoryByIdRequest) (*pb.GetCategoryByIdResponse, error) {
	category, err := s.repo.GetCategoryById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetCategoryByIdResponse{
		Category: &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
	}, nil
}
