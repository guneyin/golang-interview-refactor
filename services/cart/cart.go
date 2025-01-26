package cart

import (
	"interview/entity"
	"interview/repositories/cart"
)

type Service struct {
	repo *cart.Repository
}

func NewService() *Service {
	return &Service{repo: cart.NewRepository()}
}

func (s *Service) GetCart(sessionID string) (entity.CartItems, error) {
	return s.repo.GetCart(sessionID)
}

func (s *Service) Add(sessionID, product string, qty uint) error {
	return s.repo.AddItem(sessionID, product, qty)
}

func (s *Service) Delete(sessionID string, id uint) error {
	return s.repo.DeleteItem(sessionID, id)
}
