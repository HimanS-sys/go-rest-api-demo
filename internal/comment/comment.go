package comment

import (
	"fmt"
	"context"
)

type Store interface {
	GetComment(context.Context, string) (Comment, error) 
}

type Comment struct {
	ID string
	Slug string
	Body string
	Author string
}

// Service - all our logic resides
// on this struct.
type Service struct {
	Store Store
} 

// NewService - returns a pointer 
// to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrieving a comment")
	cmt, err := s.Store.GetComment(COmment, id)
	if err!=nil {
		fmt.Println(err)
		return Comment{}, err
	}
	return cmt, nil
}
