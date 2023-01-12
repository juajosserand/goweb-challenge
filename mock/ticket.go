package mock

import (
	"context"

	"github.com/juajosserand/goweb-challenge/internal/domain"
	"github.com/juajosserand/goweb-challenge/internal/ticket"
)

type stubRepo struct {
	db  *DbMock
	ctx context.Context
}

type DbMock struct {
	Db  []domain.Ticket
	spy bool
	err error
}

func NewRepositoryTest(dbm *DbMock) ticket.Repository {
	return &stubRepo{
		db:  dbm,
		ctx: context.Background(),
	}
}

func (r *stubRepo) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	r.db.spy = true
	if r.db.err != nil {
		return []domain.Ticket{}, r.db.err
	}
	return r.db.Db, nil
}

func (r *stubRepo) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {
	var tkts []domain.Ticket

	r.db.spy = true
	if r.db.err != nil {
		return []domain.Ticket{}, r.db.err
	}

	for _, t := range r.db.Db {
		if t.Country == destination {
			tkts = append(tkts, t)
		}
	}

	return tkts, nil
}
