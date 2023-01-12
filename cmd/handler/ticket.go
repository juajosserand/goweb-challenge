package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juajosserand/goweb-challenge/internal/ticket"
)

type Service struct {
	svc ticket.Service
}

func NewService(mux *gin.Engine, svc ticket.Service) {
	s := &Service{
		svc: svc,
	}

	ticketsMux := mux.Group("/tickets")
	ticketsMux.GET("/getByCountry/:dest", s.GetTicketsByCountry())
	ticketsMux.GET("/getAverage/:dest", s.AverageDestination())
}

func (s *Service) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {
		dest := c.Param("dest")

		total, err := s.svc.GetTotalTickets(c, dest)
		if err != nil {
			switch {
			case errors.Is(err, ticket.ErrNotFound):
				c.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
			case errors.Is(err, ticket.ErrEmptyTicketsList):
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"total_tickets": total,
		})
	}
}

func (s *Service) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {
		dest := c.Param("dest")

		avg, err := s.svc.AverageDestination(c, dest)
		if err != nil {
			switch {
			case errors.Is(err, ticket.ErrNotFound):
				c.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
			case errors.Is(err, ticket.ErrEmptyTicketsList):
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"avg": avg,
		})
	}
}
