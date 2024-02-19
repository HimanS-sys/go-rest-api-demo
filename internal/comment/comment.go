package comment

import (
	"fmt"
	"context"
	"errors"
)

var (
	ErrFetchComment = errors.New("failed to fetch comment by id.")
	ErrNotImplemented = errors.New("Not Implemented") 
)

type Comment struct {
	ID string
	Slug string
	Body string
	Author string
}

// Store - defines all the methods
// that our service needs in order to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error) 
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
		return Comment{}, ErrFetchComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
