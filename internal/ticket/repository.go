package ticket

import (
	"context"
	"errors"

	"github.com/juajosserand/goweb-challenge/internal/domain"
)

var (
	ErrEmptyTicketsList = errors.New("empty tickets list")
	ErrNotFound         = errors.New("not found")
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
}

type repository struct {
	db []domain.Ticket
}

func NewRepository(db []domain.Ticket) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	if len(r.db) == 0 {
		return []domain.Ticket{}, ErrEmptyTicketsList
	}

	return r.db, nil
}

func (r *repository) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {
	var ticketsDest []domain.Ticket

	if len(r.db) == 0 {
		return []domain.Ticket{}, ErrEmptyTicketsList
	}

	for _, t := range r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	if len(ticketsDest) == 0 {
		return []domain.Ticket{}, ErrNotFound
	}

	return ticketsDest, nil
}
