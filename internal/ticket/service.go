package ticket

import (
	"context"
)

type Service interface {
	GetTotalTickets(context.Context, string) (int, error)
	AverageDestination(context.Context, string) (float64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetTotalTickets(ctx context.Context, dest string) (int, error) {
	tickets, err := s.repo.GetTicketByDestination(ctx, dest)
	if err != nil {
		return 0, err
	}

	return len(tickets), nil
}

func (s *service) AverageDestination(ctx context.Context, dest string) (float64, error) {
	tickets, err := s.repo.GetTicketByDestination(ctx, dest)
	if err != nil {
		return 0, err
	}

	totalTickets, err := s.repo.GetAll(ctx)
	if err != nil {
		return 0, err
	}

	return 100 * float64(len(tickets)) / float64(len(totalTickets)), nil
}
